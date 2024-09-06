package main

import (
    "log"
    "net/http"
    "time"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/wailbentafat/full-stack-ecommerce/backend/internal/db"
    "github.com/wailbentafat/full-stack-ecommerce/backend/internal/auth/routes"
)
func main() {
    corsConfig := cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // Use specific origin for better security
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: false,
        MaxAge:           12 * time.Hour,
    }

    database, err := db.InitDb("database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    router := gin.Default()

    // Apply CORS middleware before defining routes
    router.Use(cors.New(corsConfig))

    // Define routes
    routes.AuthRoutes(router, database)
    routes.SecureRoutes(router)

    err = http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal(err)
    }
}
