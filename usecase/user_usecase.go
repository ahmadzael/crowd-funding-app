package controllers

import (
    "myapp/services"
    "github.com/labstack/echo/v4"
    "net/http"
)

type UserController struct {
    service services.UserService
}

func NewUserController(service services.UserService) *UserController {
    return &UserController{service: service}
}

func (c *UserController) Register(ctx echo.Context) error {
    username := ctx.FormValue("username")
    password := ctx.FormValue("password")
    if err := c.service.Register(username, password); err != nil {
        return ctx.JSON(http.StatusBadRequest, "Registration failed!")
    }
    return ctx.JSON(http.StatusOK, "User registered successfully!")
}

func (c *UserController) Login(ctx echo.Context) error {
    username := ctx.FormValue("username")
    password := ctx.FormValue("password")
    token, err := c.service.Login(username, password)
    if err != nil {
        return ctx.JSON(http.StatusUnauthorized, "Invalid credentials")
    }
    return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}
