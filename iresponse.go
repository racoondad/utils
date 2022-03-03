package utils

import "github.com/gin-gonic/gin"

type Res200 struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"ok"`
	Result interface{} `json:"result" `
	Msg    string      `json:"msg" example:"成功"`
}
type Res200List struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"ok"`
	Result ListContent `json:"result"`
	Msg    string      `json:"msg" example:"成功"`
}
type ListContent struct {
	Content  interface{} `json:"content"`
	Page     int         `json:"page" example:"1"`
	PageSize int         `json:"pageSize" example:"10"`
	Total    int64       `json:"total" example:"100"`
}
type Res400 struct {
	Code   int    `json:"code" example:"400"`
	Status string `json:"status" example:"fail"`
	Result string `json:"result" example:"错误详情"`
	Msg    string `json:"msg" example:"参数错误(400) | token无效、未认证(401) | 禁止访问(403)"`
}
type Res500 struct {
	Code   int    `json:"code" example:"500"`
	Status string `json:"status" example:"error"`
	Result string `json:"result" example:"错误详情"`
	Msg    string `json:"msg" example:"服务器错误"`
}

func ParamInvalid(c *gin.Context, err error) {
	c.JSON(400, Res400{
		Code:   400,
		Status: "fail",
		Msg:    "错误请求,请检查参数",
		Result: err.Error(),
	})
}

func Unauthorized(c *gin.Context, err error) {
	c.JSON(401, Res400{
		Code:   401,
		Status: "fail",
		Msg:    "登录失效,请重新登录",
		Result: err.Error(),
	})
}

func Forbidden(c *gin.Context, err error) {
	c.JSON(403, Res400{
		Code:   403,
		Status: "fail",
		Msg:    "权限不足,请联系管理员",
		Result: err.Error(),
	})
}
func InternalError(c *gin.Context, err error) {
	c.JSON(500, Res500{
		Code:   500,
		Status: "error",
		Msg:    "服务器内部错误",
		Result: err.Error(),
	})
}

func OK(c *gin.Context) {
	c.JSON(200, Res200{
		Code:   200,
		Status: "ok",
		Result: gin.H{},
		Msg:    "请求成功",
	})
}

func Result(c *gin.Context, v interface{}) {
	c.JSON(200, Res200{
		Code:   200,
		Status: "ok",
		Result: v,
		Msg:    "获取成功",
	})
}

func List(c *gin.Context, v interface{}, pageNo, pageSize int, total int64) {
	c.JSON(200, Res200List{
		Code:   200,
		Status: "ok",
		Result: ListContent{
			Content:  v,
			Page:     pageNo,
			PageSize: pageSize,
			Total:    total,
		},
	})
}
