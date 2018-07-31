package main

import (
    //standard
    "fmt"
    "net/http"
    "strconv"
    //github
    "github.com/gorilla/mux"
    //Original
    "./fizzbuzz"

)

var roure = mux.NewRouter()

func helloHandler(writer http.ResponseWriter, request *http.Request){
    fmt.Fprint(writer,"<h1>Hello Fizz Buzz!</h1>")
}
func fizzBuzzHandler(writer http.ResponseWriter, request *http.Request){
    vars        :=  mux.Vars(request)
    inputNum, _ :=  strconv.Atoi( vars["num"] )
    fmt.Fprint( writer,fizzbuzz.Handler( inputNum ) )
}

func httpController() {
    roure.HandleFunc("/",helloHandler)
    roure.HandleFunc("/fizzbuzz/{num:[0-9]+}",fizzBuzzHandler)
}

func main() {
    httpController()
    http.ListenAndServe(":8080", roure)
}
