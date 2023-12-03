package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 状态码
const (
	SUCCESS = 0
	FAILED  = 1
)

// Resp 接口响应
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func buildResp(c *gin.Context, resp Resp) {
	c.JSON(http.StatusOK, resp)
}

func Ok(c *gin.Context) {
	buildResp(c, Resp{SUCCESS, "success", nil})
}

func OkWithData(c *gin.Context, data interface{}) {
	buildResp(c, Resp{SUCCESS, "success", data})
}

func FailWithMsg(c *gin.Context, msg string) {
	buildResp(c, Resp{FAILED, msg, nil})
}
