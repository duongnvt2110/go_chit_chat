package main

import (
	"net/http"

	"gochitchat/routes"
)

func main(){
	
	routes.GetApplication();
	http.ListenAndServe(":8080",nil)
}