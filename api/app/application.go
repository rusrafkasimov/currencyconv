package app

import (
 "github.com/currencyconv/api/controllers"
 "github.com/gin-gonic/gin"
 "log"
 "time"
)

var router = gin.Default()

func StartApp () {
 controllers.LogInit()
 log.Print("Logger init success")
 controllers.ParseCBR()
 log.Print("Parse CBR success")
 requestmap()
 log.Print("URL map init success")
 router.Run(":8080")
 log.Print("Listen port 8080")
 go worker()
 log.Print("Worker UP")
}

func worker() {
 next := time.After(time.Minute * 30)
 for {
  select {
  case <-next:
   controllers.ParseCBR()
   next = time.After(time.Hour * 30)
  }
 }
}


