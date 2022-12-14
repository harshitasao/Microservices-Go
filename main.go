package main

import (
	"Microservice-Go/Microservices-Go/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	// "github.com/harshitasao/Microservices-go/handlers"
)

func main() {
	// NOTE: For the webserver to work in proper way we need to write and read stuff so this been acheived with the
	// help of ResponseWriter and Request.

	// here first defining where to give output and in which format
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// connecting the main with hello handler via this NewHello function
	ph := handlers.NewProducts(l)

	// creating new servemux
	// sm := http.NewServeMux()

	//
	sm := mux.NewRouter()
	// This router has a sub-router using that we can have sub-router for each methods and
	// this will provide us with more functionality with the middleware
	// Using this making router specific to GET verb and then subrouter specific for handler

	// Router for GET verb
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	// router for PUT verb
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)

	// NOTE: when we have middleware then whenevr a request comes in first it goes to router then subrouter then it will see
	// that this has a middleware so it first goes to middleware amd when it passes then goes to the subrouter

	// router for Post verb
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProducts)

	//implementing handle method of servemux type object
	// this method takes 2 parameters the path and the handler needs to be working for trhat path
	// sm.Handle("/", ph)

	// creating a basic webserver
	// and here instead of nil i am writing sm to make the default as my servemux not the defaultServerHttp
	// http.ListenAndServe(":9090", sm)

	// creating new server manually becoz the default one doesnot provide with enough functionality

	// create a new server
	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
