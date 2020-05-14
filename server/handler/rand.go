package handler

import (
	"diary/server/redis"
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

	keys, _ := redis.Keys("Diary:Content:*")
	fmt.Println("randkey", keys[0])
	content, _ := redis.Get(keys[0])
	t, _ := template.ParseFiles("template/content.html")
	fmt.Println("content", content)
	t.Execute(response, string(content))
}
