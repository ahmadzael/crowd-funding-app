package main

import (
    "myapp/config"
    "myapp/controllers"
    "myapp/repositories"
    "myapp/services"
    "github.com/labstack/echo/v4"
)

func main() {
    config.InitDB()
    e := echo.New()

    campaignRepo := repositories.NewCampaignRepository(config.DB)
    campaignService := services.NewCampaignService(campaignRepo)
    campaignController := controllers.NewCampaignController(campaignService)

    e.POST("/campaign", campaignController.AddCampaign)
    e.PUT("/campaign/:id/approve", campaignController.ApproveCampaign)
    e.POST("/campaign/:id/donate", campaignController.DonateToCampaign)

    e.Start(":8081")
}
