package main

import (
    "crowdfunding/config"
    "crowdfunding/controllers"
    "crowdfunding/repositories"
    "crowdfunding/services"
    "github.com/labstack/echo/v4"
)

func main() {
    config.InitDB()
    e := echo.New()

    userRepo := repositories.NewUserRepository(config.DB)
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

    e.POST("/register", userController.Register)
    e.POST("/login", userController.Login)

    e.Start(":8080")
}
