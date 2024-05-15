package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leotang5451/bitoTest/middleware"
)

func RegisterApi(r *gin.Engine) {
	r.Use(middleware.Base())

	//單身匹配
	sp := r.Group("/singlePerson")
	sp.GET("/querySinglePeople", GetSinglePeopleHandler)                //查詢單身匹配
	sp.POST("/addSinglePersonAndMatch", AddSinglePersonAndMatchHandler) //新增單身匹配
	sp.DELETE("/removeSinglePerson", RemoveSinglePersonHandler)         //移除單身匹配
}
