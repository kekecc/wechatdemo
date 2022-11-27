package post

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
)

func GetUserImage(userid uint) string {
	var user model.User
	db := database.Get()
	if err := db.Where("id = ?", userid).First(&user).Error; err != nil {
		log.Println("获取用户信息失败", err)
		return ""
	}
	return user.Fileid
}
