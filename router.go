package main

import (
	"github.com/gorilla/mux"
	"yugo/controller/web"
)

func InitRouter(router *mux.Router) {

	router.HandleFunc("/hello", web.Get)

	// 可以用一个middleware函数包裹controller，进行过滤，通过就执行controller

}
