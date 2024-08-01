package main

import "net/http"

func handlehello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func main(){
     server:=http.NewServeMux()
	http.HandleFunc("/hello",handlehello)
   
	fs :=http.FileServer(http.Dir("./public"))

	err:=http.ListenAndServe(":8080",server)
    if err!=nil{
		panic(err)
	}

}