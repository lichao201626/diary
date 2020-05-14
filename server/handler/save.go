package handler

import (
	"diary/server/redis"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type SaveBody struct {
	username string
	password string
	content  string
}

func SaveHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Handler Hello")
	fmt.Println("r.Method = ", request.Method)
	fmt.Println("r.URL = ", request.URL)
	fmt.Println("r.Header = ", request.Header)

	fmt.Println("r.Body = ", request.Body)

	var req map[string]interface{}
	body, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(body, &req)

	fmt.Printf("POST form-urlencoded: username=%s, password=%s\n", body, req)

	key := "Diary:Content:" + req["username"].(string) + req["password"].(string)
	redis.Set(key, []byte(req["content"].(string)))

	t, _ := template.ParseFiles("template/content.html")
	t.Execute(response, req["content"].(string))
}
