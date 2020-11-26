package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.HandlerFunc(makeTopPage))
	http.ListenAndServe(":8888", nil)
}

func makeTopPage(resw http.ResponseWriter, req *http.Request) {
	content, err := ioutil.ReadFile("top.html.txt")
	if err != nil {
		fmt.Println("テンプレートファイル読み込み失敗", err)
		os.Exit(1)
	}
	templateString := string(content)
	templ := template.Must(template.New("sendname").Parse(templateString))
	templ.Execute(resw, req.FormValue("yourname"))
}
