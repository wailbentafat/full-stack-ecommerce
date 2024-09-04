package main

import (
    "log"
    "net/http"
    "github.com/wailbentafat/full-stack-ecommerce/backend/internal/db"
)

func main() {
    database, err := db.InitDb("database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
