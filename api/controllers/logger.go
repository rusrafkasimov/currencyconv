package controllers

import (
	"log"
	"os"
)

var infoLog = log.New(LogFile(), "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(LogFile(), "ERROR\t", log.Ldate|log.Ltime)

func LogInit () {
	initl, err := os.Create("Application.log")
	if err != nil {
		log.Fatal(err)
	}
	defer initl.Close()
}

func LogFile () *os.File {
	f, err := os.OpenFile("Application.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return f
}