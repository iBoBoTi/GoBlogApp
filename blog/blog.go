package blog

type Post struct {
	Id		string
	Title	string
	Body	string
}

type Posts struct{
	Posts []Post
}


var GolangBlog Posts = Posts{Posts: []Post{}}
