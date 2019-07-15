package config

type Config struct {
	Domain  string
	Session *Session
	Render  *Render
}

type Session struct {
	Secret string
}

type Render struct {
	IsDevelopment bool
}
