package handler

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leotang5451/bitoTest/request"
	"github.com/leotang5451/bitoTest/response"
	"github.com/shopspring/decimal"
)

// 使用者
type SinglePerson struct {
	UserId         int             `json:"user_id"`          //id
	Name           string          `json:"name"`             //姓名
	Height         decimal.Decimal `json:"height"`           //身高
	Gender         string          `json:"gender"`           //性別 0:男 1:女
	NumberWantDate int             `json:"number_want_date"` //人數
	IsDelete       bool            `json:"-"`                //是否刪除
}

// 使用者列表
var SinglePersonList []SinglePerson

// 新增匹配
func AddSinglePersonAndMatchHandler(c *gin.Context) {
	req := request.AddSinglePersonAndMatchReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	data := SinglePerson{
		Name:           req.Name,
		Height:         req.Height,
		Gender:         req.Gender,
		NumberWantDate: req.NumberWantDate,
	}

	if len(SinglePersonList) == 0 {
		data.UserId = 1
		SinglePersonList = append(SinglePersonList, data)

		resp := response.BaseResponse{
			Code: response.MatchFail,
		}

		c.JSON(http.StatusOK, resp)
		return
	}

	//匹配使用者
	possibleMatches, code, err := findMatches(&data)
	if err != nil {
		c.Error(err)
		return
	}

	//排序
	sort.Slice(SinglePersonList, func(i, j int) bool {
		return SinglePersonList[i].UserId < SinglePersonList[j].UserId
	})

	//給user_id
	data.UserId = len(SinglePersonList) + 1
	SinglePersonList = append(SinglePersonList, data)

	resp := response.BaseResponse{
		Code: code,
		Data: func() interface{} {
			if possibleMatches != nil {
				return possibleMatches
			}
			return nil
		}(),
	}

	c.JSON(http.StatusOK, resp)
}

// 查詢匹配
func GetSinglePeopleHandler(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	result := []SinglePerson{}
	for _, v := range SinglePersonList {
		if !v.IsDelete && v.NumberWantDate > 0 {
			result = append(result, v)
			limit--
			if limit == 0 {
				break
			}
		}
	}

	resp := response.BaseResponse{
		Code: response.Success,
		Data: result,
	}

	c.JSON(http.StatusOK, resp)
}

// 移除匹配
func RemoveSinglePersonHandler(c *gin.Context) {
	req := request.RemoveSinglePersonReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	for v := range SinglePersonList {
		if SinglePersonList[v].UserId == req.UserId {
			SinglePersonList[v].IsDelete = true
			break
		}
	}

	resp := response.BaseResponse{
		Code: response.Success,
	}

	c.JSON(http.StatusOK, resp)
}

// 匹配使用者
func findMatches(person *SinglePerson) (result *SinglePerson, code string, err error) {
	for v := range SinglePersonList {
		switch person.Gender {
		//男生
		case "0": //男身高>女身高
			sub := person.Height.Sub(SinglePersonList[v].Height)
			if !SinglePersonList[v].IsDelete && SinglePersonList[v].Gender == "1" &&
				SinglePersonList[v].NumberWantDate > 0 && sub.Cmp(decimal.NewFromInt(0)) > 0 {
				code = response.MatchSuccess
				SinglePersonList[v].NumberWantDate--
				person.NumberWantDate--
			} else {
				code = response.MatchFail
			}
			//女生
		case "1": //男身高>女身高
			sub := person.Height.Sub(SinglePersonList[v].Height)
			if !SinglePersonList[v].IsDelete && SinglePersonList[v].Gender == "0" &&
				SinglePersonList[v].NumberWantDate > 0 && sub.Cmp(decimal.NewFromInt(0)) <= 0 {
				code = response.MatchSuccess
				SinglePersonList[v].NumberWantDate--
				person.NumberWantDate--
			} else {
				code = response.MatchFail
			}
		}

		//用完約會次數
		if SinglePersonList[v].NumberWantDate == 0 {
			SinglePersonList[v].IsDelete = true
		}
	}

	return
}
