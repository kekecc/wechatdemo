package handler

import (
	databaseuser "wechatdemo/database/user"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetUserImage(c *gin.Context) {
	userid := c.GetUint("user")
	fileid := databaseuser.GetUserImage(userid)
	if fileid == "" {
		response.Failed(c, 400, "未获取到用户fileid", "")
		return
	}
	response.Success(c, 200, "成功获取用户fileid", "")
}
