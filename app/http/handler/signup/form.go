package signup

import "net/http"

func (u *User) Form(w http.ResponseWriter, r *http.Request) {
	if err := u.re.Form(w); err != nil {
		panic(err)
	}
}
