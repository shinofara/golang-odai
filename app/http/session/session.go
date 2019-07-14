package session

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"golang-odai/domain/user"
	"net/http"
)

const SESSION_NAME = "login_session"

type Session struct {
	store *sessions.CookieStore
}

func init() {
	gob.Register(&user.User{})
}

// New returns new Session
func New(secret string) *Session {
	//本来はハードコーディングせずに外部から渡すこと
	store := sessions.NewCookieStore([]byte(secret))

	store.Options = &sessions.Options{
		Domain:     "localhost",
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
func (sess *Session) SetUser(w http.ResponseWriter, r *http.Request, user *user.User) error {
	session, err := sess.store.Get(r, SESSION_NAME)
	if err != nil {
		return err
	}
	session.Values["user"] = user

	return session.Save(r, w)
}

// GetUser gets user session
func (sess *Session) GetUser(r *http.Request) (*user.User, error) {
	session, err := sess.store.Get(r, SESSION_NAME)
	if err != nil {
		return nil, err
	}
	u, ok := session.Values["user"].(*user.User)
	if !ok {
		return nil, errors.New("user session not found")
	}

	if u == nil {
		return nil, errors.New("user session not found")
	}

	return u, nil
}
