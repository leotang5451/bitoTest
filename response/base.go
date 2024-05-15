package response

const (
	Success      string = "1"  //成功
	MatchSuccess string = "2"  //匹配成功
	Fail         string = "-1" //錯誤
	MatchFail    string = "-2" //匹配失敗
)

type BaseResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Code  string      `json:"code,omitempty"`
	Error string      `json:"err,omitempty"`
}
