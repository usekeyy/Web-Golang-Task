package main

import (
	route "Golang-Task/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFile("/assets/css/add.css", "./assets/css/add.css")
	router.StaticFile("/assets/css/base.css", "./assets/css/base.css")
	router.StaticFile("/assets/css/index.css", "./assets/css/index.css")

	router.LoadHTMLGlob("views/*")

	route.Routes(router)

	router.Run()
}
