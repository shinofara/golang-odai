package main

import (
	"context"
	"golang-odai/adapter/http/render"
	"golang-odai/adapter/http/session"
	"golang-odai/config"
	"golang-odai/external/http/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
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
	}

	r, err := route.New(cfg)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    ":80",
		Handler: r,
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
