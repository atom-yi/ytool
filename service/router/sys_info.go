package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/user"
)

// UsernameRouter 获取用户名
func UsernameRouter(c *gin.Context) {
	userInfo, err := user.Current()
	if err != nil {
		c.String(http.StatusOK, "unknown user")
		return
	}

	c.String(http.StatusOK, userInfo.Username)
}
