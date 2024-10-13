package main

import (
	"context"
	"fmt"
	"gst-billing/commons/constants"
	"gst-billing/router"
	"gst-billing/utils/postgres"
	"log"
)

// @title GST Billing API
// @version 1.0
// @description This is a sample API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @x-extention-openapi {"example": "This is a sample API"}
// @host localhost:8080
// @BasePath /
func main() {
	ctx := context.Background()
	// initialize postgres client
	err := postgres.InitPostgresDBConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer postgres.ClosePostgres(ctx)

	startRouter()
}

func startRouter() {
	router := router.GetRouter()
	fmt.Printf(constants.ServerRunningMessage, constants.PortDefaultValue)
	err := router.Run(fmt.Sprintf(":%d", constants.PortDefaultValue))
	if err != nil {
		fmt.Println(err)
	}
}
