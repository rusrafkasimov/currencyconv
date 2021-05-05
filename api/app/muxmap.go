package app

import "github.com/currencyconv/api/controllers"

func requestmap() {
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.IndexHome)
	router.POST("/", controllers.Index)
}
