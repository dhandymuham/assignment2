package main

import (
	"test/config"
	"test/controllers"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "test/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	port = ":9000"
)

func main() {
	db := config.InitDB()
	controllerOrders := &controllers.InDB{DB: db}
	router := gin.Default()

	router.GET("/orders", controllerOrders.GetOrders)

	router.POST("/orders", controllerOrders.CreateOrder)

	router.PUT("/orders/:orderId", controllerOrders.UpdateOrder)

	router.DELETE("/orders/:orderId", controllerOrders.DeleteOrder)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(port)
}
