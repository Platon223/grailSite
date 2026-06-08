package main

import (
    "log"
    "net/http"
    "github.com/Platon223/grailSite/internal/routes"
)

func main() {
    mux := http.NewServeMux()
    routes.Register(mux)

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
