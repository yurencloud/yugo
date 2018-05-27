package web

import (
	"net/http"
	"yugo/view"
	"github.com/CloudyKit/jet"
)

func Get(writer http.ResponseWriter, request *http.Request) {
	//log.Println("user controller")
	////log.Println(request.PostFormValue("name"))
	////log.Println(request.Form["name"][0])
	//log.Println(request.Form)
	//
	//vars := mux.Vars(request)
	//age := vars["age"]
	//log.Println(age)

	varMap := make(jet.VarMap)
	varMap.Set("name", "tom")

	view.Render(writer, "index.html", varMap)
}
