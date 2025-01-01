package main

import(
	"fmt"
	"net/http"
	"html/template"
)

var templates *template.Template

func main() {
  mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
  mux.Handle("/static/",http.StripPrefix("/static/",fs))
  
	mux.HandleFunc("/site",sitehandler)
  mux.HandleFunc("/form",formhandler)
	mux.HandleFunc("/boot",boothabdler)
	mux.HandleFunc("/today",todayhandler)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080",mux)
}

func init() {
	templates = template.Must(template.ParseGlob("web/templates/*.html"))
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html",nil)
	if err != nil{
		fmt.Println("Template execution error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func sitehandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "site.html",nil)
	if err != nil{
		fmt.Println("Template execution error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func boothabdler(w http.ResponseWriter,r *http.Request){
	err := templates.ExecuteTemplate(w,"boot.html",nil)
	if err!=nil{
    fmt.Println("テンプレートエラー")
		http.Error(w,"インターネットエラー",http.StatusInternalServerError)
	}
}

func todayhandler(w http.ResponseWriter,r *http.Request){
	err := templates.ExecuteTemplate(w,"calendar.html",nil)
	if err!=nil{
		fmt.Println("テンプレートエラー")
		http.Error(w,"インターネットエラー",http.StatusInternalServerError)
	}
}