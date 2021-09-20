package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Welcome To My Blog</h1>")
}

func main(){
	http.HandleFunc("/", HomeFunc)
	err:=http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
