package main

import (
    "log"
    "net/http"
    "fmt"
)

func main() {
    router := CreateRouter()
    fmt.Println("Server listining on port: 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}