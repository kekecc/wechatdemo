package handler

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"
	"wechatdemo/utils"

	"github.com/gin-gonic/gin"
)

func UpdateUserAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		response.Failed(c, 400, "获取头像错误", nil)
		return
	}
	f, _ := file.Open()
	if err = utils.Upload("avatar", file.Filename, f); err != nil {
		log.Println(err)
		response.Failed(c, 400, "上传头像错误", nil)
		return
	}
	f.Close()

	userid := c.GetUint("user")
	db := database.Get()
	if userid != 0 {
		err := db.Model(&model.User{}).Where("id = ?", userid).Update("fileid", file.Filename).Error
		if err != nil {
			log.Println("更改失败", err)
			response.Failed(c, 400, "更改失败", err)
			return
		}
		response.Success(c, 200, "更改成功", file.Filename)
	}
}

func UpdateUserDetail(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Println("参数绑定出现问题")
		response.Failed(c, 400, "参数绑定出现问题", err)
		return
	}
	if json["name"] == nil && json["qq"] == nil && json["wx"] == nil {
		log.Println("参数错误", json)
		response.Failed(c, 400, "参数错误", nil)
		return
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
