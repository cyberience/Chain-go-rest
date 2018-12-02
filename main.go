package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "github.com/AENCO-Global/Chain-go-sdk"
)

func main() {
    route := mux.NewRouter()

    route.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Some Variables" ,r)
    })

    route.HandleFunc("/version" , func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Version" ,":1")
    })


    // Set up the listerner for the incoming requests
    http.ListenAndServe(":3001", route)
}