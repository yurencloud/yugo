package yugo

import (
	"github.com/gorilla/mux"
	"yugo/config"
	"net/http"
	"strings"
	r "yugo/router"
)

func staticServer(router *mux.Router) {
	staticConfig := config.Get("static")
	staticArray := strings.Split(staticConfig,",")
	// 生成一个或多个静态目录，默认static,可自行修改，或添加，以英文逗号分隔
	for index := range staticArray {
		static := staticArray[index]
		router.PathPrefix("/").Handler(http.StripPrefix("/"+static, http.FileServer(http.Dir(static))))
	}
}

func Run()  {



	router := mux.NewRouter()

	r.InitRouter(router)

	staticServer(router)


	port := config.Get("port")

	//log.Printf("服务器正在运行，端口：%s",port)
	http.ListenAndServe(":"+port, router)
}