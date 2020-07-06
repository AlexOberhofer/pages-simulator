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

func handleBlog(w http.ResponseWriter, r *http.Request) {

    http.ServeFile(w, r, "blog/blog.html")
    fmt.Printf("Called Blog Handler")
}

func main() {

	fmt.Printf("Server started on port 8081. Press Ctrl + C to exit.\n")

    fs := http.FileServer(http.Dir("public"))

    http.Handle("/public/", http.StripPrefix("/public/", fs)) 

    http.HandleFunc("/", handleHome)

    http.HandleFunc("/blog", handleBlog)

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
