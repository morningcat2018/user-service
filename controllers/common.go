package controllers

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func buildSuccessRes(data interface{}) *ResponseData {
	return &ResponseData{0, "success", data}
}
