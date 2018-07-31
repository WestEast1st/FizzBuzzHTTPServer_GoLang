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

/*
*　controller
 */
func helloController(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprint(writer, "<h1>Hello Fizz Buzz!</h1>")
}

func fizzBuzzController(writer http.ResponseWriter, request *http.Request) {
    vars := mux.Vars(request)
	inputNum, err := strconv.Atoi(vars["num"])
	if err != nil {
        StatusBadRequest(writer)
        return
    }
	fmt.Fprint(writer, "<h1>"+fizzbuzz.Handler(inputNum)+"</h1>")
}

/*
* error
 */
func StatusBadRequest(writer http.ResponseWriter) {
    writer.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(writer, "<h1>invalid number</h1>")
}

/*
* HTTP Route
 */
func httpRoute() {
    roure.HandleFunc("/", helloController)
    roure.HandleFunc("/fizzbuzz/{num}", fizzBuzzController)
}

/*
* Main
* Main -> HTTP Route -> controller -> FizzBuzz
*
 */
func main() {
    httpRoute()
    http.ListenAndServe(":8080", roure)
}
