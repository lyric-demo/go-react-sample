package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	SUCCESS       = 0
	UNKNOWN       = 1
	PARSEERROR    = 98
	SESSIONEXPIRE = 99
)

// ResResult 响应数据
type ResResult struct {
	Status  int         `json:"status"` // 响应状态码
	Data    interface{} `json:"data"`   // 响应数据
	Message string      `json:"msg"`    // 响应消息
}

// BaseController 提供基础操作
type BaseController struct {
	beego.Controller
}

// ParseBody 解析请求的body数据
func (this *BaseController) ParseBody(v interface{}) error {
	return json.Unmarshal(this.Ctx.Input.RequestBody, v)
}

// ResParseError 响应请求数据解析错误
func (this *BaseController) ResParseError(err error) {
	this.Error(PARSEERROR, "请求的数据格式不正确！", err)
}

// ResJson 响应json数据
func (this *BaseController) ResJson(v interface{}) {
	this.Data["json"] = v
	this.ServeJSON()
}

// Success 响应成功数据
func (this *BaseController) Success(data interface{}, msg string) {
	result := ResResult{
		Status:  SUCCESS,
		Data:    data,
		Message: msg,
	}
	this.ResJson(result)
}

// Error 响应异常数据
func (this *BaseController) Error(status int, msg string, err error) {
	result := ResResult{
		Status:  status,
		Message: msg,
	}
	beego.Error(fmt.Sprintf("%s:%v", msg, err))
	this.ResJson(result)
}
