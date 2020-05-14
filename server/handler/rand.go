package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type body struct {
	username string
	password string
}

func RandHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Handler Hello")
	fmt.Println("r.Method = ", request.Method)
	fmt.Println("r.URL = ", request.URL)
	fmt.Println("r.Header = ", request.Header)

	fmt.Println("r.Body = ", request.Body)

	// randKey, _ := redis.RandomKey()
	// fmt.Println("randkey", randKey)
	// content, _ := redis.Get(string(randKey))
	t, _ := template.ParseFiles("template/content.html")
	t.Execute(response, string(""))
}
