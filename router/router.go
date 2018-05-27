package router

import (
	"github.com/gorilla/mux"
	"yugo/controller/web"
)

func InitRouter(router *mux.Router) {

	router.HandleFunc("/hello", web.Get)

}
