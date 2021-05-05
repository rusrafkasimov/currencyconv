package controllers

import (
	"github.com/currencyconv/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHome (*gin.Context) {
	return
}

func Index (c *gin.Context) {
	var req models.ReqResCurrency
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid_JSON"})
		errorLog.Printf("Invalid_JSON %v \n", err)
		return
	}
    if !CurrencyValid(req.FrCur, req.ToCur) {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid_Currency"})
		errorLog.Printf("Invalid_Currency \n")
	} else {
		req.ToVal = Convertation(req.FrCur, req.ToCur, req.FrVal)
		c.JSON(200, req)
		infoLog.Printf("Successful_Response %v \n", req)
	}
}
