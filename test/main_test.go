package test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leotang5451/bitoTest/handler"
)

// test Gin server
var testR *gin.Engine

func TestMain(m *testing.M) {
	testR = gin.Default()
	handler.RegisterApi(testR)

	os.Exit(m.Run())
}

func resetData() {
	handler.SinglePersonList = nil
}
