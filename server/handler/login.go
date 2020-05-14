package handler

import (
	"diary/server/redis"
	"fmt"
	"html/template"
	"net/http"
)

// LoginHandler first wirte content page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)

	r.ParseForm()

	// 第一种方式
	// username := request.Form["username"][0]
	// password := request.Form["password"][0]

	// 第二种方式
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	fmt.Printf("POST form-urlencoded: username=%s, password=%s\n", username, password)

	key := "Diary:User:" + username + "-" + password
	redis.Set(key, []byte("online"))

	t, _ := template.ParseFiles("template/content.html")

	t.Execute(w, "")
}
