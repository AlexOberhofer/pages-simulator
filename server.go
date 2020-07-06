package main

import (
    "fmt"
    "log"
    "net/http"
)

var home_counter int = 0

func handleHome(w http.ResponseWriter, r *http.Request) {
    home_counter++
    http.ServeFile(w, r, "index.html")
    fmt.Printf("Home Handler called. Count: %d\n", home_counter)
}

func handleGNU(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "gnuboy/gnuboy.html")
    fmt.Printf("Called Blog Handler")
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Called Download Handler")
}

func handleProject(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "projects/projects.html")
	fmt.Printf("Called Project Handler")
}

func startServer() {
	fmt.Printf("Attempting to start web server...")
}


func main() {

	fmt.Printf("Server started on port 8080. Press Ctrl + C to exit.\n")

    fs := http.FileServer(http.Dir("public"))

    http.Handle("/public/", http.StripPrefix("/public/", fs)) 

    http.HandleFunc("/", handleHome)

    http.HandleFunc("/gnuboy", handleGNU)

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}