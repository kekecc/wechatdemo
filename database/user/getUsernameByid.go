package post

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
)

func GetUserNameByID(userid uint) (string, error) {
	var user model.User
	db := database.Get()
	err := db.Where("id=?", userid).Find(&user).Error
	if err != nil {
		log.Println("由id查找本名出错")
		return "", err
	}
	//log.Println("userid ", userid, " 对应 ", user.Name)
	return user.Name, nil
}
func GetUserqqAndWxByID(postid uint) (string, string, error) {
	var user model.User
	var post model.Post
	db := database.Get()
	err := db.Where("id = ?", postid).Find(&post).Error
	if err != nil {
		log.Println("由id查找帖子出错")
		return "", "", nil
	}
	err = db.Where("id = ?", post.UserId).Find(&user).Error
	if err != nil {
		log.Println("由id查找本人出错")
		return "", "", nil
	}
	//log.Println("userid ", post.UserId, " 对应 ", user.Name)
	return user.QQ, user.Wx, err
}

func GetUserDetailById(id uint) (string, string, string) {
	db := database.Get()
	var user model.User
	if err := db.Model(&model.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		log.Println("获取用户信息错误", err.Error())
		return "", "", ""
	}
	return user.Fileid, user.QQ, user.Wx
}

