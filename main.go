package main

import (
	"context"
	"github.com/pkg/errors"
	"go.opencensus.io/plugin/ochttp"
	"golang-odai/adapter/http/render"
	"golang-odai/adapter/http/session"
	"golang-odai/config"
	"golang-odai/external/firebase"
	"golang-odai/external/http/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"

)

func main() {
	cfg := &config.Config{
		Domain: "localhost",
		Session: &session.Config{
			Domain: "localhost",
			Secret: "xxxxx",
		},
		Render: &render.Config{
			IsDevelopment: true,
		},
		Firebase: &firebase.Config{
			ApiKEY: os.Getenv("FIREBASE_API_KEY"),
		},
		Jaeger: &jaeger.Options{
			AgentEndpoint:     "trace:6831",
			CollectorEndpoint: "http://trace:14268/api/traces",
			ServiceName:       "golang-odai",
		},
	}

	r, err := route.New(cfg)
	if err != nil {
		panic(err)
	}

	// add tracer
	if err := tracer(cfg); err != nil {
		panic(err)
	}

	och := &ochttp.Handler{
		Handler: r,
		GetStartOptions: func(r *http.Request) trace.StartOptions {
			startOptions := trace.StartOptions{}
			if r.URL.Path == "/healthcheck" {
				startOptions.Sampler = trace.NeverSample()
			}
			return startOptions
		},
	}

	srv := &http.Server{
		Addr:    ":80",
		Handler: och,
	}

	// Graceful Shutdown
	// SIGTERM発火後動作中のプロセスを即時停止させるのではなく、受付を停止して、処理中のリクエストがなくなるまで起動を続ける。
	// すべてのリクエストの処理が完了したらHTTPサーバを停止させる。
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Fatalln("Server closed with error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Println("Failed to gracefully shutdown:", err)
	}
	log.Println("Server shutdown")
}

func tracer(cfg *config.Config) error {
	ex, err := jaeger.NewExporter(*cfg.Jaeger)
	if err != nil {
		return errors.Wrap(err, "failed to create the Jaeger exporter")
	}
	trace.RegisterExporter(ex)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	return nil
}