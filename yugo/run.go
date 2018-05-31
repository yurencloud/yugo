package yugo

import (
	"github.com/gorilla/mux"
	"yugo/config"
	"net/http"
	"strings"
	"github.com/gorilla/csrf"
	"strconv"
)

func staticServer(router *mux.Router) {
	staticConfig := config.Get("static")
	staticArray := strings.Split(staticConfig, ",")
	// 生成一个或多个静态目录，默认static,可自行修改，或添加，以英文逗号分隔
	for index := range staticArray {
		static := staticArray[index]
		router.PathPrefix("/").Handler(http.StripPrefix("/"+static, http.FileServer(http.Dir(static))))
	}
}

func Run() {

	router := mux.NewRouter()

	InitRouter(router)

	staticServer(router)

	configMap := config.GetConfigMap()

	maxAge, _ := strconv.Atoi(configMap["csrf.max.age"])

	CSRF := csrf.Protect(
		[]byte(configMap["csrf.key"]),
		csrf.RequestHeader(configMap["csrf.request.header"]),
		csrf.FieldName(configMap["csrf.field.name"]),
		csrf.MaxAge(maxAge),
	)

	http.ListenAndServe(":"+configMap["port"], CSRF(router))
}
