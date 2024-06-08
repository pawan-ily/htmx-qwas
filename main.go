package main

import (
	"html/template"
	"net/http"
	"log"
)

type Todo struct {
	Id      int
	Message string
}
func main(){

	data := map[string][]Todo{
		"Todos": {
			{Id: 1, Message: "hire me"},
		},}
	todoHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
	
		templ.Execute(w, data)
	}

	addtodohandler := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
		data["Todos"] = append(data["Todos"], todo)
 
		templ := template.Must(template.ParseFiles("index.html"))
		templ.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", todoHandler)
	http.HandleFunc("/add-todo", addtodohandler)

	log.Fatal(http.ListenAndServe("8000", nil))
}
