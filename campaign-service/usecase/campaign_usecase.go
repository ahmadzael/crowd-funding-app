package controllers

import (
    "myapp/services"
    "myapp/utils"
    "github.com/labstack/echo/v4"
    "net/http"
)

type CampaignController struct {
    service services.CampaignService
}

func NewCampaignController(service services.CampaignService) *CampaignController {
    return &CampaignController{service: service}
}

func (c *CampaignController) AddCampaign(ctx echo.Context) error {
    tokenString := ctx.Request().Header.Get("Authorization")

    // Call gRPC Auth Service to validate token
    valid, message := utils.ValidateTokenWithAuthService(tokenString)
    if !valid {
        return ctx.JSON(http.StatusUnauthorized, message)
    }

    // Proceed with campaign creation
    // ...

    return ctx.JSON(http.StatusOK, "Campaign created successfully!")
}

func (c *CampaignController) ApproveCampaign(ctx echo.Context) error {
    tokenString := ctx.Request().Header.Get("Authorization")

    // Call gRPC Auth Service to validate token
    valid, message := utils.ValidateTokenWithAuthService(tokenString)
    if !valid {
        return ctx.JSON(http.StatusUnauthorized, message)
    }

    // Proceed with campaign approval
    // ...

    return ctx.JSON(http.StatusOK, "Campaign status updated successfully!")
}

func (c *CampaignController) DonateToCampaign(ctx echo.Context) error {
    tokenString := ctx.Request().Header.Get("Authorization")

    // Call gRPC Auth Service to validate token
    valid, message := utils.ValidateTokenWithAuthService(tokenString)
    if !valid {
        return ctx.JSON(http.StatusUnauthorized, message)
    }

    // Proceed with donation
    // ...

    return ctx.JSON(http.StatusOK, "Donation successful!")
}
