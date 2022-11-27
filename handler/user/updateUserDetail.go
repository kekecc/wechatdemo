package handler

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func UpdateUserDetail(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Println("参数绑定出现问题")
		response.Failed(c, 400, "参数绑定出现问题", err)
		return
	}
	if json["name"] == nil && json["qq"] == nil && json["wx"] == nil && json["fileid"] == nil {
		log.Println("参数错误", json)
		response.Failed(c, 400, "参数错误", nil)
		return
	}
	if json["fileid"] != nil {
		log.Println(json["fileid"])
	}
	userid := c.GetUint("user")
	db := database.Get()
	if userid != 0 {
		err := db.Model(&model.User{}).Where("id = ?", userid).Updates(json).Error
		if err != nil {
			log.Println("更改失败", err)
			response.Failed(c, 400, "更改失败", err)
			return
		}
		response.Success(c, 200, "更改成功", json)
	}
}
