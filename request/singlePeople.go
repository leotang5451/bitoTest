package request

import "github.com/shopspring/decimal"

// 新增使用者
type AddSinglePersonAndMatchReq struct {
	Name           string          `json:"name"`             //姓名
	Height         decimal.Decimal `json:"height"`           //身高
	Gender         string          `json:"gender"`           //性別 0:男 1:女
	NumberWantDate int             `json:"number_want_date"` //人數
}

// 移除使用者
type RemoveSinglePersonReq struct {
	UserId int `json:"user_id"`
}
