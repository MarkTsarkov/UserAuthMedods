package main

import (
    "os"

    middleware "github.com/marktsarkov/userauthmedods/middleware"
    routes "github.com/marktsarkov/userauthmedods/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        port = "8000"
    }

    router := gin.New()
    router.Use(gin.Logger())
    routes.UserRoutes(router)

    router.Use(middleware.Authentication())

    router.Run(":" + port)
}