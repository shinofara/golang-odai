package model

import (
	"context"
)

type Post struct {
	ID uint32
	Name string
	Text string
}

func Select(ctx context.Context) ([]Post, error) {
	db, err := New()
	if err != nil {
		return nil, err
	}
	rows, err := db.Open().QueryContext(
		ctx,
		"select id, name, text from posts",
	)
	if err != nil {
		return nil, err
	}

	list := make([]Post, 0)
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Name, &p.Text); err != nil {
			return nil, err
		}
		list = append(list, p)
	}

	return list, nil
}

func Insert(ctx context.Context, post Post) error {
	db, err := New()
	if err != nil {
		return err
	}
	if _, err := db.Open().ExecContext(
		ctx,
		"insert into posts (name, text) values (?,?)",
		post.Name,
		post.Text,
	); err != nil {
		return err
	}

	return nil
}