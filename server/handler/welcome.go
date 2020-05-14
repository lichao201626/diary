package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

// WelcomeHandler first page
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)

	t, _ := template.ParseFiles("template/login.html")

	t.Execute(w, "")
}
