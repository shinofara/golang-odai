package config

import (
	"golang-odai/adapter/http/render"
	"golang-odai/adapter/http/session"
)

type Config struct {
	Domain  string
	Session *session.Config
	Render  *render.Config
}