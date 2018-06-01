package app

import (
	// orderController "github.com/audio35444/tp2-api/src/app/controllers/order"
	itemController "github.com/audio35444/tp2-api/src/app/controllers/item"
	pingController "github.com/audio35444/tp2-api/src/app/controllers/ping"
	tokenController "github.com/audio35444/tp2-api/src/app/controllers/token"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	router.GET("/ping", pingController.Ping)
	router.POST("/tokens", tokenController.AddToken)
	router.PUT("/tokens/:token", tokenController.ResetToken)
	router.GET("/tokens/:docId", tokenController.GetToken)
	router.GET("/items/:docId", itemController.GetItem)
	router.DELETE("/items/:docId", itemController.DeleteItem)
	router.PUT("/items", itemController.UpdateItem)
	router.POST("/items", itemController.AddItem)
	router.GET("/items", itemController.GetItems)
	router.GET("/", func(c *gin.Context) {
		c.JSON(400, gin.H{
			"error": "Not found",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080  router.Run("":8000"")
}

func exampleFunc(valores ...int) {
	//viene un slice con los datos
}
