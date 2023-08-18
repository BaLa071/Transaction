package controllers

import (
	"net/http"
	"strings"
	"transaction/interfaces"
	"transaction/models"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService interfaces.ITransaction
}

func InitTranasctionController(transactionService interfaces.ITransaction) TransactionController {
	return TransactionController{transactionService}
}

func (pc *TransactionController) CreateTransaction(ctx *gin.Context) {
	var profile *models.Transaction
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := pc.TransactionService.CreateTransaction(profile)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}

func (pc *TransactionController) GetTransactionById(ctx *gin.Context) {
}
