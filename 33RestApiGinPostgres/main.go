package main

import (
	"learnapi/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.MyRouter(r)

	log.Fatal(r.Run())

}
