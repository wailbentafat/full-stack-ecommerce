package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/wailbentafat/full-stack-ecommerce/backend/internal/db"
    "github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/routes"
   
)

func main() {

    database, err := db.InitDb("database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

  
    router := gin.Default()


    routes.AuthRoutes(router, database)
    routes.SecureRoutes(router)
   
    err = http.ListenAndServe(":8080", router)

    if err != nil {
        log.Fatal(err)
    }
}
