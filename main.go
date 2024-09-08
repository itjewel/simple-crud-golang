package main

import (
    "fmt"
    "log"
    "net/http"
    "simple-crud-golang/handlers"
)

func main() {
    handlers.InitTemplates()  // Initialize templates here

    http.HandleFunc("/", homePage) // Define the home page route
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.HandleFunc("/create", handlers.CreatePage)
    http.HandleFunc("/read", handlers.ReadPage)
    http.HandleFunc("/update", handlers.UpdatePage)
    http.HandleFunc("/delete", handlers.DeletePage)

    fmt.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// homePage handler definition
func homePage(w http.ResponseWriter, r *http.Request) {
    // You can render a home page template or send plain text for now
    fmt.Fprint(w, "Welcome to the Simple CRUD Homepage")
}
