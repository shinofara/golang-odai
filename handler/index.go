package handler

import (
	"net/http"
	"golang-odai/model"
	"html/template"
	"github.com/go-chi/chi"
)

type Data struct{
	Posts []model.Post
}

func IndexRender(w http.ResponseWriter,posts []model.Post) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("template/index.html"))
	data := Data{
		posts,
	}

	// テンプレートを描画
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := model.Select(r.Context())
	if err != nil {
		/*
		if err == model.Notfound {
			not found
		}
		*/

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	IndexRender(w, posts)
}

func PostDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	posts, err := model.FindByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	IndexRender(w, posts)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("template/form.html"))
	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	text := r.FormValue("text")

	p := model.Post{
		Name: name,
		Text: text,
	}

	if err := model.Insert(r.Context(), p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}