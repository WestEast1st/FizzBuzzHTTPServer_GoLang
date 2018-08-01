package main

import (
	//standard
	"fmt"
	"net/http"
	"strconv"
	//github
	"github.com/gorilla/mux"
    "github.com/WestEast1st/FizzBuzzHTTPServer_GoLang/modules"
)

var route = mux.NewRouter()

/*
*　controller
*   helloController : top
*   fizzBuzzController : fizzBuzz
 */
func helloController(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Hello Fizz Buzz!</h1>")
}

func fizzBuzzController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	inputNum, err := strconv.Atoi(vars["num"])
	//Processing other than int.
	if err != nil {
		StatusBadRequest(writer)
		return
	}
	//success
	fmt.Fprint(writer, "<h1>"+modules.Handler(inputNum)+"</h1>")
}

/*
* error
*   StatusBadRequest : 400 invalid number
 */
func StatusBadRequest(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(writer, "<h1>invalid number</h1>")
}

/*
* HTTP Route
 */
func httpRoute() {
	route.HandleFunc("/", helloController)
	route.HandleFunc("/fizzbuzz/{num}", fizzBuzzController)
}

/*
* Main
* Main -> HTTP Route -> controller -> FizzBuzz
*
 */
func main() {
	httpRoute()
	http.ListenAndServe(":8080", route)
}
