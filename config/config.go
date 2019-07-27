package config

import (
	"contrib.go.opencensus.io/exporter/jaeger"
	"golang-odai/adapter/http/render"
	"golang-odai/adapter/http/session"
	"golang-odai/external/firebase"
)

type Config struct {
	Domain  string
	Session *session.Config
	Render  *render.Config
	Firebase *firebase.Config
	Jaeger *jaeger.Options
}