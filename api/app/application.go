package app

import (
 "github.com/currencyconv/api/controllers"
 "github.com/gin-gonic/gin"
 "log"
)

var router = gin.Default()

func StartApp () {
 controllers.LogInit()
 log.Print("Logger init success")
 controllers.ParseCBR()
 log.Print("Parse CBR success")
 requestmap()
 router.Run(":8080")
}



