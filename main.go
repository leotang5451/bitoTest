package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leotang5451/bitoTest/handler"
)

func main() {
	r := gin.Default()
	handler.RegisterApi(r)

	r.Run(`:8080`)
}
