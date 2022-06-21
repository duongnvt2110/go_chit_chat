package home

import (
	"fmt"
	"net/http"
)


func Index(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Hello World");
}
