package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func WriteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)

	t, _ := template.ParseFiles("template/content.html")

	t.Execute(w, "")
}
