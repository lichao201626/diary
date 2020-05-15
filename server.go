package main

import (
	"diary/server/handler"
	"fmt"
	"net/http"
)

func main() {
	//监听协议
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.HandleFunc("/", handler.WelcomeHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/write", handler.WriteHandler)
	http.HandleFunc("/rand", handler.RandHandler)
	http.HandleFunc("/save", handler.SaveHandler)

	//监听服务
	err := http.ListenAndServe("0.0.0.0:8880", nil)

	if err != nil {
		fmt.Println("服务器错误", err)
	}
}
