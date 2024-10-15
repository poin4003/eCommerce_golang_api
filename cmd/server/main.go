package main

import (
	_ "github.com/poin4003/eCommerce_golang_api/cmd/swag/docs"
	"github.com/poin4003/eCommerce_golang_api/internal/initalize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Documentation Ecommerce Backend SHOPDEVGO
// @version 1.0.0
// @description This is a sample server celler server
// @termsOfService github.com/poin4003/eCommerce_golang_api

// @contact.name TEAM pchuy
// @contact.url github.com/poin4003/eCommerce_golang_api
// @contact.email pchuy4003@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /v1/2024
// @schema http

func main() {
	r := initalize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8000")
}
