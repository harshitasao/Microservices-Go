package main

import (
	"log"
	"net/http"
	"os"

	"github.com/harshitasao/Microservices-go/handlers"
)

func main() {
	// NOTE: For the webserver to work in proper way we need to write and read stuff so this been acheived with the
	// help of ResponseWriter and Request.

	// here first defining where to give output and in which format
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// connecting the main with hello handler via this NewHello function
	hh := handlers.NewHello(l)

	// creating new servemux
	sm := http.NewServeMux()
	//implementing handle method of servemux type object
	// this method takes 2 parameters the path and the handler needs to be working for trhat path
	sm.Handle("/", hh)

	// creating a basic webserver
	// and here instead of nil i am writing sm to make the default as my servemux not the defaultServerHttp
	http.ListenAndServe(":9090", sm)
}
