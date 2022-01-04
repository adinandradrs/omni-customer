package response

type BaseResponse struct {
	Message string                 `json:"message"`
	Result  bool                   `json:"result"`
	Data    map[string]interface{} `json:"data"`
}
