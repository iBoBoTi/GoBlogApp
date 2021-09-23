package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/iBoBoTi/BlogApp/blog"
	"html/template"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"

)


func index(rw http.ResponseWriter, req *http.Request) {
	// handler for home/root page as well as display all post
	header:= rw.Header()
	header.Add("Content-Type","text/html")
	rw.WriteHeader(200)

	data:= blog.GolangBlog
	tmpl, _ := template.ParseFiles("templates/index.html")
	err := tmpl.Execute(rw, data)
	if err != nil {
		return 
	}
}

func PostCreate(rw http.ResponseWriter, req *http.Request){
	// handler to present a form to add post to the user
	tmpl, _ := template.ParseFiles("templates/create.html")
	err := tmpl.Execute(rw, nil)
	if err != nil {
		return 
	}
}

func PostUpdate(rw http.ResponseWriter, req *http.Request){
	//id := strings.TrimSpace(chi.URLParam(req,""))

}

func PostDetail(rw http.ResponseWriter, req *http.Request){
	// handler to retrieve a post and display it to the client
	id := chi.URLParam(req,"Id")
	var data blog.Post

	for _,v:=range blog.GolangBlog.Posts{
		if id == v.Id{
			data = v
		}
	}
	tmpl, _ := template.ParseFiles("templates/post_detail.html")
	err := tmpl.Execute(rw,data)
	if err != nil {
		return 
	}
}

func PostDelete(rw http.ResponseWriter, req *http.Request){
	id := chi.URLParam(req,"Id")
	for i,v:=range blog.GolangBlog.Posts{
		if id == v.Id{
			blog.GolangBlog.Posts = append(blog.GolangBlog.Posts[:i],blog.GolangBlog.Posts[i+1:]...)
		}
	}
	http.Redirect(rw,req,"/",301)
}

func AddFormHandler(rw http.ResponseWriter, req *http.Request){
	// handler take the create post form and add to GolangBlog
	req.ParseForm()

	if req.FormValue("Title") != "" && req.FormValue("Body") != ""{
		post := blog.Post{
			Id: uuid.NewString(),
			Title: req.FormValue("Title"),
			Body: req.FormValue("Body"),
		}
		blog.GolangBlog.Posts = append(blog.GolangBlog.Posts,post)
	}

	tmpl, _ := template.ParseFiles("templates/post_confirm.html")
	err := tmpl.Execute(rw, nil)
	if err != nil {
		return 
	}
}


func main() {
	r := chi.NewRouter()

	// middleware
	r.Use(middleware.Logger)

	r.Get("/",index)
	r.Get("/create",PostCreate)
	r.Post("/add", AddFormHandler)
	r.Put("/post/{Id}", PostUpdate)
	r.Get("/{Id}", PostDetail)
	r.Get("/post/{Id}", PostDelete)

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
