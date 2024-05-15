package handler

import (
	"github.com/gin-gonic/gin"
	docs "github.com/leotang5451/bitoTest/docs"
	"github.com/leotang5451/bitoTest/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterApi(r *gin.Engine) {
	r.Use(middleware.Base())
	docs.SwaggerInfo.BasePath = ""

	//單身匹配
	sp := r.Group("/singlePerson")
	sp.GET("/querySinglePeople", GetSinglePeopleHandler)                //查詢單身匹配
	sp.POST("/addSinglePersonAndMatch", AddSinglePersonAndMatchHandler) //新增單身匹配
	sp.DELETE("/removeSinglePerson", RemoveSinglePersonHandler)         //移除單身匹配

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
