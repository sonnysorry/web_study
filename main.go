package main

import (
	"net/http"
	"log"
	"fmt"
)
var Port = ":8055"
func main() {
	http.HandleFunc("/",ServeFiles)
	fmt.Println("Serving at : ","127.0.0.1" + Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}

func ServeFiles(w http.ResponseWriter, r *http.Request){

	path := r.URL.Path

	fmt.Println(path)

	if path == "/"{
		path = "./assets/login.html"
	}else{
		path = "."+path
	}
	http.ServeFile(w,r,path)
}