package controllers

import (
	"encoding/json"
	"github.com/currencyconv/api/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func IndexAPI (*gin.Context) {
	return
}

func ParseCBR() {
	// Connection
	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	var req = http.Client{
		Timeout: time.Second * 120,
	}

	// Get JSON CBR
	res, err := req.Get(url)
	if err != nil {
		errorLog.Printf("Parse_CBR_JSON")
	}
	defer res.Body.Close()

	// Read JSON CBR
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		errorLog.Printf("Read_JSON")
	}

	// Unmarshall data to structure
	var data models.JsonCBR
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		errorLog.Printf("Unmarshalling")
	}

	CurData = data
	infoLog.Printf("CBR_parse_success")
}

func Index (c *gin.Context) {
	var req models.ReqResCurrency
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid_JSON"})
		errorLog.Printf("Invalid_JSON %S", err)
	}
	req.ToVal = Convertation(req.FrCur, req.ToCur, req.FrVal)
	c.JSON(200, req)
	infoLog.Printf("Successful_response")
}