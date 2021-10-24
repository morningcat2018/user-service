package controllers

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func buildSuccessRes(data interface{}) *ResponseData {
	return &ResponseData{0, "success", data}
}

func buildFailRes(e error) *ResponseData {
	return &ResponseData{-1, e.Error(), nil}
}
