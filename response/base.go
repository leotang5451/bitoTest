package response

const (
	Success      string = "1"  //成功
	MatchSuccess string = "2"  //匹配成功
	Fail         string = "-1" //錯誤
	MatchFail    string = "-2" //匹配失敗
)

type BaseResponse struct {
	//匹配資料
	Data interface{} `json:"data,omitempty"`
	//1:成功 2:匹配成功 -1:錯誤 -2:匹配失敗
	Code string `json:"code,omitempty"`
	//錯誤訊息
	Error string `json:"err,omitempty"`
}
