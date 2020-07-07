package main

import (
    "fmt"
    "log"
    "net/http"
)

var home_counter int = 0

/**
* Handler function for routing to the home page
*/
func handleHome(w http.ResponseWriter, r *http.Request) {
    home_counter++
    http.ServeFile(w, r, "index.html")
    fmt.Printf("Home Handler called. Count: %d\n", home_counter)
}

func startServer() {
    fmt.Printf("Attempting to start web server...\n")
    
    //Initialize File Server
    http.Handle("/", http.FileServer(http.Dir("./")))

    //add a home handler
    http.HandleFunc("./", handleHome)

    //add ping handler
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Pong!")
    })

    fmt.Printf("Server initialized successfully!\n")

}

func main() {

    startServer()

    fmt.Printf("Server started on port 8080. Press Ctrl + C to exit.\n")

    log.Fatal(http.ListenAndServe(":8080", nil))

}