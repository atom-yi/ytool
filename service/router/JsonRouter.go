package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"ytool/common"
)

type JsonReqParam struct {
	Indent  string `json:"indent"`
	Content string `json:"content"`
}

// JsonRouter format JSON string
func JsonRouter(c *gin.Context) {
	var reqParam JsonReqParam
	if err := c.ShouldBindJSON(&reqParam); err != nil {
		common.FailWithMsg(c, err.Error())
		return
	}

	var data interface{}
	if err := json.Unmarshal([]byte(reqParam.Content), &data); err != nil {
		common.FailWithMsg(c, err.Error())
		return
	}

	var r []byte
	var err error
	if reqParam.Indent == "" {
		r, err = json.Marshal(data)
	} else {
		r, err = json.MarshalIndent(data, "", reqParam.Indent)
	}

	if err != nil {
		common.FailWithMsg(c, err.Error())
		return
	}

	common.OkWithData(c, string(r))
}
