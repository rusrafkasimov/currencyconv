package app

import "github.com/currencyconv/api/controllers"

func requestmap() {
	router.GET("/", controllers.IndexAPI)
	router.POST("/", controllers.Index)
}
