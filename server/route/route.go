package route

import (
	"net/http"
)

type Router struct {
	Path string
	Handler function
}

type route struct {}

var routes = []Router{

}

func (rt *route) new() {
	for router := range routes {
		http.HandlerFunc(router.Path, router.Handler)
	}
}