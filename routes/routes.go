package routes

import (
	"gochitchat/controllers/home"
	"gochitchat/controllers/users"
	"net/http"
)


func GetApplication(){
	http.HandleFunc("/", home.Index)
	//users 
	// http.HandleFunc("/users", users.Index)
	http.HandleFunc("/users/create", users.Create)
	http.HandleFunc("/users", users.Index)
	http.HandleFunc("/users/update", users.Update)
	http.HandleFunc("/users/delete", users.Delete)

	// http.HandleFunc("/posts", posts.Index)
	// http.HandleFunc("/threads", threads.Index)
	// http.HandleFunc("/sessions", sessions.Index)
}



