package main

import (
	"fmt"
	"net/http"
	"html/template"
)
func sayHello(w http.ResponseWriter, r *http.Request){
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
		if err !=nil {
			fmt.Println("template read err",err)	
			return
		}

	//渲染模板
	err = t.Execute(w,"hahahahahah")
		if err != nil {
			fmt.Println("render template ",err)
			return
		}
}
func main(){
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":9000",nil)	
	if err !=nil {
		fmt.Println("Http start fail, err:", err)
		return
	}
}



