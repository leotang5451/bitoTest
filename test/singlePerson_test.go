package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/leotang5451/bitoTest/handler"
	"github.com/leotang5451/bitoTest/request"
	"github.com/leotang5451/bitoTest/response"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var (
	addSinglePersonAndMatch    = `/singlePerson/addSinglePersonAndMatch`
	removeSinglePersonAndMatch = `/singlePerson/removeSinglePerson`
	querySinglePersonAndMatch  = `/singlePerson/querySinglePeople`
)

// 新增匹配人
func addPerson(data request.AddSinglePersonAndMatchReq) (result response.BaseResponse, err error) {
	reqByte, _ := json.Marshal(data)
	req, err := http.NewRequest(http.MethodPost, addSinglePersonAndMatch, bytes.NewBuffer(reqByte))
	if err != nil {
		return
	}

	recorder := httptest.NewRecorder()
	testR.ServeHTTP(recorder, req)

	err = json.Unmarshal(recorder.Body.Bytes(), &result)
	return
}

// 新增男生 170 沒有匹配
func TestAddSinglePersonManNoMatch(t *testing.T) {
	resetData()

	height, _ := decimal.NewFromString("170")
	reqData := request.AddSinglePersonAndMatchReq{
		Name:           "男生1",
		Height:         height,
		Gender:         "0",
		NumberWantDate: 1,
	}

	result, err := addPerson(reqData)
	if assert.NoError(t, err) {
		assert.Equal(t, response.MatchFail, result.Code)
	}
}

// 新增女生 169 沒有匹配
func TestAddSinglePersonWomanNoMatch(t *testing.T) {
	resetData()

	height, _ := decimal.NewFromString("170")
	reqData := request.AddSinglePersonAndMatchReq{
		Name:           "女生1",
		Height:         height,
		Gender:         "1",
		NumberWantDate: 1,
	}

	result, err := addPerson(reqData)
	if assert.NoError(t, err) {
		assert.Equal(t, response.MatchFail, result.Code)
	}
}

// 新增女生 169 有匹配
func TestAddSinglePersonMatch(t *testing.T) {
	resetData()

	//新增男生 170 沒有匹配
	TestAddSinglePersonManNoMatch(t)

	height, _ := decimal.NewFromString("169")
	reqData := request.AddSinglePersonAndMatchReq{
		Name:           "女生1",
		Height:         height,
		Gender:         "1",
		NumberWantDate: 1,
	}

	result, err := addPerson(reqData)
	if assert.NoError(t, err) {
		assert.Equal(t, response.MatchSuccess, result.Code)
	}
}

// 新增女生 170 有匹配
func TestAddSinglePersonNoMatch(t *testing.T) {
	//新增男生 170 沒有匹配
	TestAddSinglePersonManNoMatch(t)

	height, _ := decimal.NewFromString("170")
	reqData := request.AddSinglePersonAndMatchReq{
		Name:           "女生1",
		Height:         height,
		Gender:         "1",
		NumberWantDate: 1,
	}

	result, err := addPerson(reqData)
	if assert.NoError(t, err) {
		assert.Equal(t, response.MatchSuccess, result.Code)
	}
}

// 新增女生 169 有匹配、新增女生 168 沒有匹配
func TestAddSinglePersonWomanMatchAndNoMatch(t *testing.T) {
	//新增男生 170 沒有匹配
	TestAddSinglePersonManNoMatch(t)

	height, _ := decimal.NewFromString("169")
	reqData := request.AddSinglePersonAndMatchReq{
		Name:           "女生1",
		Height:         height,
		Gender:         "1",
		NumberWantDate: 1,
	}

	result, err := addPerson(reqData)
	if assert.NoError(t, err) {
		assert.Equal(t, response.MatchSuccess, result.Code)
	}

	height, _ = decimal.NewFromString("168")
	reqData2 := request.AddSinglePersonAndMatchReq{
		Name:   "女生2",
		Height: height,
		Gender: "1",
	}

	result, err = addPerson(reqData2)
	if assert.NoError(t, err) {
		assert.Equal(t, response.MatchFail, result.Code)
	}
}

// 新增男生 169 沒有匹配、新增男生 170 有匹配
func TestAddSinglePersonManMatchAndNoMatch(t *testing.T) {
	//新增女生 169 沒有匹配
	TestAddSinglePersonWomanNoMatch(t)

	height, _ := decimal.NewFromString("169")
	reqData := request.AddSinglePersonAndMatchReq{
		Name:           "男生1",
		Height:         height,
		Gender:         "1",
		NumberWantDate: 1,
	}

	result, err := addPerson(reqData)
	if assert.NoError(t, err) {
		assert.Equal(t, response.MatchFail, result.Code)
	}

	height, _ = decimal.NewFromString("168")
	reqData2 := request.AddSinglePersonAndMatchReq{
		Name:   "女生2",
		Height: height,
		Gender: "1",
	}

	result, err = addPerson(reqData2)
	if assert.NoError(t, err) {
		assert.Equal(t, response.MatchFail, result.Code)
	}
}

// 刪除匹配
func TestRemoveSinglePerson(t *testing.T) {
	resetData()

	//新增3個人
	height, _ := decimal.NewFromString("170")
	reqData := request.AddSinglePersonAndMatchReq{
		Name:           "男生1",
		Height:         height,
		Gender:         "0",
		NumberWantDate: 1,
	}
	addPerson(reqData)

	reqData = request.AddSinglePersonAndMatchReq{
		Name:           "男生2",
		Height:         height,
		Gender:         "0",
		NumberWantDate: 1,
	}
	addPerson(reqData)

	reqData = request.AddSinglePersonAndMatchReq{
		Name:           "男生3",
		Height:         height,
		Gender:         "0",
		NumberWantDate: 1,
	}
	addPerson(reqData)

	// 移除男生2
	delReqData := request.RemoveSinglePersonReq{
		UserId: 2,
	}
	reqByte, _ := json.Marshal(delReqData)
	req, err := http.NewRequest(http.MethodDelete, removeSinglePersonAndMatch, bytes.NewBuffer(reqByte))
	if err != nil {
		return
	}

	recorder := httptest.NewRecorder()
	testR.ServeHTTP(recorder, req)

	result := response.BaseResponse{}
	err = json.Unmarshal(recorder.Body.Bytes(), &result)
	if assert.NoError(t, err) {
		assert.Equal(t, response.Success, result.Code)
	}
}

// 查詢匹配
func TestQuerySinglePerson(t *testing.T) {
	resetData()

	//新增3個人
	height, _ := decimal.NewFromString("170")
	reqData := request.AddSinglePersonAndMatchReq{
		Name:           "男生1",
		Height:         height,
		Gender:         "0",
		NumberWantDate: 1,
	}
	addPerson(reqData)

	height, _ = decimal.NewFromString("171")
	reqData = request.AddSinglePersonAndMatchReq{
		Name:           "男生2",
		Height:         height,
		Gender:         "0",
		NumberWantDate: 1,
	}
	addPerson(reqData)

	height, _ = decimal.NewFromString("172")
	reqData = request.AddSinglePersonAndMatchReq{
		Name:           "男生3",
		Height:         height,
		Gender:         "0",
		NumberWantDate: 1,
	}
	addPerson(reqData)

	// 查詢2個可以匹配
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(`%s?limit=%s`, querySinglePersonAndMatch, "2"), nil)
	if err != nil {
		return
	}

	recorder := httptest.NewRecorder()
	testR.ServeHTTP(recorder, req)

	result := response.BaseResponse{}
	err = json.Unmarshal(recorder.Body.Bytes(), &result)
	if assert.NoError(t, err) {
		assert.Equal(t, response.Success, result.Code)

		res := (result.Data.([]interface{}))

		checkData := []handler.SinglePerson{}
		for _, v := range res {
			val := reflect.ValueOf(v)
			if val.Kind() != reflect.Map {
				continue
			}

			data := val.Interface().(map[string]interface{})

			height, _ := decimal.NewFromString(data["height"].(string))
			checkData = append(checkData, handler.SinglePerson{
				Name:   data["name"].(string),
				Height: height,
				Gender: data["gender"].(string),
			})
		}

		assert.Equal(t, "男生1", checkData[0].Name)
		assert.Equal(t, "170", checkData[0].Height.String())
		assert.Equal(t, "0", checkData[0].Gender)

		assert.Equal(t, "男生2", checkData[1].Name)
		assert.Equal(t, "171", checkData[1].Height.String())
		assert.Equal(t, "0", checkData[1].Gender)
	}
}
