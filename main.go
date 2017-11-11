package main

import "html/template"
import "net/http"
import "log"
import "time"


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	log.Printf("start server port 8011")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/",ind)
	http.ListenAndServe(":8011", nil)
	
	track(time.Now(),"test")
}
func ind(w http.ResponseWriter, Req *http.Request) {
	tpl.ExecuteTemplate(w,"index.gohtml",nil)
}
func track(start time.Time ,name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}