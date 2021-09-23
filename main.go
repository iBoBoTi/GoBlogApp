package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/iBoBoTi/BlogApp/blog"
	"html/template"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5/middleware"
	//"github.com/go-chi/chi"
)


func index(rw http.ResponseWriter, req *http.Request) {
	// handler for home/root page
	header:= rw.Header()
	header.Add("Content-Type","text/html")
	rw.WriteHeader(200)

	data:= blog.GolangBlog
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(rw, data)
}

func PostCreate(rw http.ResponseWriter, req *http.Request){
	tmpl, _ := template.ParseFiles("templates/create.html")
	tmpl.Execute(rw, nil)
}

func PostUpdate(rw http.ResponseWriter, req *http.Request){}

func PostDetail(rw http.ResponseWriter, req *http.Request){}

func PostDelete(rw http.ResponseWriter, req *http.Request){}

func AddFormHandler(rw http.ResponseWriter, req *http.Request){
	req.ParseForm()
	if req.Method != http.MethodPost{
		return
	}
	post := blog.Post{
		Title: req.FormValue("Title"),
		Body: req.FormValue("Body"),
	}
	blog.GolangBlog.Posts = append(blog.GolangBlog.Posts,post)

	tmpl, _ := template.ParseFiles("templates/post_confirm.html")
	tmpl.Execute(rw, nil)
}


func main() {
	r := chi.NewRouter()

	// middleware
	r.Use(middleware.Logger)

	r.Get("/",index)
	r.Get("/create",PostCreate)
	r.Post("/add", AddFormHandler)
	http.HandleFunc("/edit", PostUpdate)
	http.HandleFunc("/detail", PostDetail)
	http.HandleFunc("/delete", PostDelete)

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
