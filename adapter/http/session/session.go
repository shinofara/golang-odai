package session

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"golang-odai/domain"
	"net/http"
)

const SESSION_NAME = "login_session"

type Session struct {
	store *sessions.CookieStore
}

func init() {
	gob.Register(&domain.User{})
}

// New returns new Session
func New(domain, secret string) *Session {
	//本来はハードコーディングせずに外部から渡すこと
	//Production,Developmentといった動作環境に応じて値を変更する事
	store := sessions.NewCookieStore([]byte(secret))

	store.Options = &sessions.Options{
		Domain:     domain,
		Path:       "/",
		MaxAge:     3600,
		Secure:     false,
		HttpOnly:   true,
	}

	return &Session{
		store: store,
	}
}

// SetUser sets user session
func (sess *Session) SetUser(w http.ResponseWriter, r *http.Request, user *domain.User) error {
	session, err := sess.store.Get(r, SESSION_NAME)
	if err != nil {
		return err
	}
	session.Values["user"] = user

	return session.Save(r, w)
}

// GetUser gets user session
func (sess *Session) GetUser(r *http.Request) (*domain.User, error) {
	session, err := sess.store.Get(r, SESSION_NAME)
	if err != nil {
		return nil, err
	}
	u, ok := session.Values["user"].(*domain.User)
	if !ok {
		return nil, errors.New("user session not found")
	}

	if u == nil {
		return nil, errors.New("user session not found")
	}

	return u, nil
}