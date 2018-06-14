package view

import (
	"github.com/CloudyKit/jet"
	"bytes"
	"net/http"
	"yugo/config"
	_ "yugo/log"
	"github.com/sirupsen/logrus"
)


func Render(writer http.ResponseWriter, templateName string, varMap jet.VarMap) {
	View := GetInstance()
	template, err := View.GetTemplate(templateName + config.Get("template.suffix"))
	if err != nil {
		logrus.Error("模板文件不存在")
	}
	var html bytes.Buffer // needs to conform to io.Writer interface (like gin's context.Writer for example)
	if err = template.Execute(&html, varMap, nil); err != nil {
		logrus.Error("模板文件渲染出错")
	}
	writer.Write(html.Bytes())
}

// 把Render分开, 以便添加global方法
// 添加全局方法，建议先在当前项目实现GetInstance->global->RenderView=>Render方法，以后就只调用当前项目的Render
// 添加局部方法，就临时使用GetInstance->global->RenderView
func GetInstance() *jet.Set {
	var View = jet.NewHTMLSet(config.Get("template.path"))
	// global在这里添加

	return View
}

func RenderView(writer http.ResponseWriter, templateName string, varMap jet.VarMap, View *jet.Set)  {
	template, err := View.GetTemplate(templateName + config.Get("template.suffix"))
	if err != nil {
		logrus.Error("模板文件不存在")
	}
	var html bytes.Buffer // needs to conform to io.Writer interface (like gin's context.Writer for example)
	if err = template.Execute(&html, varMap, nil); err != nil {
		logrus.Error("模板文件渲染出错")
	}
	writer.Write(html.Bytes())
}
