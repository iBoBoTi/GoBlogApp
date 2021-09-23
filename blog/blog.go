package blog

type Post struct {
	Id		string
	Title	string
	Body	string
}

type BlogPost struct{
	Posts []Post
}


var Post1 Post = Post{Id: "01",Title: "How To Make A Billion GoRoutines", Body: "Start small go big is the only way you can do this."}

var Post2 Post = Post{Id: "02",Title: "A Thousand Ways To Die Well", Body: "Why do you want to die?"}

var GolangBlog BlogPost = BlogPost{Posts: []Post{Post1,Post2}}
