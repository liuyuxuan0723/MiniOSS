package handler

import "net/http"

type Response struct {
	Code int
	Msg  string
	Data map[string]interface{}
}

func (r *Response) OK() *Response {
	r.Code = http.StatusOK
	r.Msg = "success"
	return r
}

func (r *Response) ERROR() *Response {
	r.Code = http.StatusInternalServerError
	r.Msg = "server error"
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) SetMsg(msg string) *Response {
	r.Msg = msg
	return r
}

func (r *Response) SetData(key string, data interface{}) *Response {
	if r.Data == nil {
		r.Data = make(map[string]interface{})
	}
	r.Data[key] = data
	return r
}
