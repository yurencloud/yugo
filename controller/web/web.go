package web

import (
	"net/http"
	"yugo/view"
	"github.com/CloudyKit/jet"
	"yugo/session"
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

	s := session.GetInstance(writer, request)

	//s.Values["name"] = "tom1"
	//s.Save(request, writer)

	varMap := make(jet.VarMap)
	varMap.Set("name", s.Values["name"])


	view.Render(writer, "index", varMap)
}
