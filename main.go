package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed static
	dirStatic embed.FS
)

func main() {
	http.HandleFunc("/", tplHandler)

	fs := http.FileServer(http.FS(dirStatic))
	http.Handle("/static/", fs)
	_ = http.ListenAndServe(":80", nil)
}

func tplHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Panicln("加载模板错误", err)
		return
	}

	if err = tpl.Execute(w, nil); err != nil {
		log.Panicln("模板渲染错误", err)
		return
	}
}
