package routes

import (
	"transaction/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoute(router *gin.Engine, controller controllers.TransactionController) {
	router.POST("/api/Transaction/create", controller.CreateTransaction)
	router.GET("/api/Transaction/:id", controller.GetTransactionById)
}
