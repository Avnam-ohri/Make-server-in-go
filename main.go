package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Thank you\n")
	name:=r.FormValue("name")
	Address:=r.FormValue("address")
	fmt.Fprintf(w,"Name= %s\n", name)
	fmt.Fprintf(w,"Address= %s\n", Address)
	
}




func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w , "Method is not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}


func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

fmt.Printf("Server has begun\n")
if err:= http.ListenAndServe(":8080",nil); err !=nil {
	log.Fatal(err)
}	
}

