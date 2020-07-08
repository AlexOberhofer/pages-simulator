///////////////////////////////////////////////////////////////////////////////
//
// A simple web server
//
// License: GNU GPLv3
// (C) 2020 Alex Oberhofer
///////////////////////////////////////////////////////////////////////////////
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var verbose = 1

/**
* Handler function for routing to the home page
 */
func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
	fmt.Printf("Home handler called...\n")
}

func logStdOut(logMessage string) {
	if verbose > 0 {
		fmt.Printf(logMessage)
		fmt.Printf("\n")
	}
}

/**
* Start server and assign fileservers and handlers
 */
func startServer() {
	logStdOut("Attempting to start web server...")

	//Initialize File Server
	http.Handle("/", http.FileServer(http.Dir("./")))

	//add a home handler
	http.HandleFunc("./", handleHome)

	//add ping handler
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!")
	})

	logStdOut("Server initialized successfully!")

}

/*
* Print usage message to console
 */
func usage() {
	fmt.Printf("go run server.go \n")
	fmt.Printf("Flags: \n")
	fmt.Printf("p <port> -v <log level>\n")
}

func main() {

	port := flag.String("p", "8080", "web server port")
	//TODO: pull verbose flag here
	flag.Parse()

	startServer()

	fmt.Printf("Server started on port %s. Press Ctrl + C to exit.\n", *port)

	log.Fatal(http.ListenAndServe(":"+*port, nil))

}
