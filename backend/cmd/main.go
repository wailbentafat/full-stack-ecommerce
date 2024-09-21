package main

import (
    "log"
    "net/http"
    "time"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/wailbentafat/full-stack-ecommerce/backend/internal/db"
    "github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/routes"
    "github.com/wailbentafat/full-stack-ecommerce/backend/internal/product/routes"
    "github.com/wailbentafat/full-stack-ecommerce/backend/internal/core/cach"
    )
func main() {
    
    corsConfig := cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: false,
        MaxAge:           12 * time.Hour,
    }
    log.Println("Server started on port 8080")
    fileCache, err := cach.NewFilecach("cache")
    if err != nil {
        log.Fatal(err)
    }

    database, err := db.InitDb("database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    router := gin.Default()

    router.Use(cors.New(corsConfig))

    routes.AuthRoutes(router, database)
    routes.SecureRoutes(router)
    product_routes.Routes(router, database, fileCache)

    err = http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal(err)
    }
}
