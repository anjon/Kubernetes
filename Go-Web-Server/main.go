package main

// Import Packages
import (
    "log"
    "net/http"
)

func main() {

    // Server the Desired HTML File
    http.Handle("/", http.FileServer(http.Dir("./content")))

    log.Fatal(http.ListenAndServe(":9091", nil))
}
