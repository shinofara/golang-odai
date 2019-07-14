package post

import "golang-odai/domain/user"

type Post struct {
	ID uint32
	UserID uint32
	Text string

	User *user.User
}