package main

import(
	"fmt"
	"net/http"
)


func main(){
	
	http.HandleFunc("/", Test)
	println("Listening at :8080 please wait for a while")
	http.ListenAndServe(":8080",nil)
}


func Test(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello World "+"amd/64 Go1.14.1")
}