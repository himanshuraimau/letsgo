package main

import (
	"net/http"
	"text/template"
)

func handlehello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request){
	html,err:=template.ParseFiles("templates/index.tmpl")
    
	if err!=nil{
	    w.Write([]byte("Error parsing template")) 
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	html.Execute(w,"TEST DATA") 
}


func main(){
     server:=http.NewServeMux()
	http.HandleFunc("/hello",handlehello)
   
	fs :=http.FileServer(http.Dir("./public"))
	server.Handle("/",fs)

	err:=http.ListenAndServe(":8080",server)
    if err!=nil{
		panic(err)
	}

}