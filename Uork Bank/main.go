package main

import (
    "fmt"
    "log"
    "net/http"

    "uorkbank/handlers"
)

func main() {
    http.HandleFunc("/", handlers.SubscriptionHandler)

    port := ":2308"
    fmt.Printf("Servidor rodando em http://localhost%s\n =)", port)

    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatal(err)
    }
}
