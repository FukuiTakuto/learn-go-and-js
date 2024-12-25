package main

import(
	"fmt"
	"net/http"
	"html/template"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("../web/templates/*.html"))
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html",nil)
	if err != nil{
		fmt.Println("Template execution error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	fs := http.FileServer(http.Dir("../web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/",homehandler)
  fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}