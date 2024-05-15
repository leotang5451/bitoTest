package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leotang5451/bitoTest/response"
)

func Base() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("recover panic%v\n", err)
			}
		}()

		c.Next()

		//error
		if len(c.Errors) > 0 {
			lastError := c.Errors.Last()
			resp := response.BaseResponse{
				Code:  response.Fail,
				Error: lastError.Error(),
			}

			c.JSON(http.StatusBadRequest, resp)
			return
		}
	}
}
