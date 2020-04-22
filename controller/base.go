package controller

import "github.com/kataras/iris/sessions"

type Response struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

type Lists struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

func ApiResource(status bool, objects interface{}, msg string) (r *Response) {
	r = &Response{Status: status, Data: objects, Msg: msg}
	return
}

type Responses struct {
	Status int        `json:"status"`
	Success    interface{} `json:"success"`
	Message   interface{} `json:"message"`
}

func Api(status int, success interface{}, message string) (r *Responses) {
	r = &Responses{Status: status, Success: success, Message: message}
	return
}

type Resp struct {
	Status int        `json:"status"`
	Data    interface{} `json:"data"`
}
func ApiResp(status int, data interface{}) (r *Resp) {
	r = &Resp{Status: status, Data:data}
	return
}

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	Sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
