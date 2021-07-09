package main

import ("fmt"
        "net/http")

func home_page(page http.ResponseWriter, r *http.Request){
  fmt.Fprintf(page, "Practice is very hard!")
}

func contacts_page(page http.ResponseWriter, r *http.Request){
  fmt.Fprintf(page, "first page!")
}

func handleRequest(){
  //fmt.Println("Hello world")
  http.HandleFunc("/", home_page)//home_page это название метода
  http.HandleFunc("/firstpage/", contacts_page) //contacts просто название, может быть любым
  http.ListenAndServe(":8080", nil)//Порт, nil это null или none(настройки для самого сервера)
}

func main() {
  handleRequest()
}
