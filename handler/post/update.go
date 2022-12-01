package post

import (
	"encoding/json"
	"log"
	"reflect"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	db := database.Get()
	userId := c.GetUint("user")
	var user model.User
	log.Println("当前正在更新的人", userId)
	if userId == 0 {
		response.Failed(c, 400, "当前用户不存在token", err)
		return
	}
	db.Where("id = ?", userId).First(&user)
	jsons := make(map[string]interface{})
	if err := c.BindJSON(&jsons); err != nil {
		response.Failed(c, 400, "给定更新参数错误!", err)
		return
	}
	var post model.Post
	log.Println("postid:", jsons["postid"])
	if jsons["postid"] == 0 || jsons["postid"] == nil {
		log.Println("postid为0")
		response.Failed(c, 400, "postid不能为0", nil)
		return
	} else {
		log.Println("postid为", jsons["postid"], " 类型为", reflect.TypeOf(jsons["postid"]))
	}
	if err := db.Where("id = ?", jsons["postid"]).First(&post).Error; err != nil {
		log.Println("未成功查询到帖子记录", err)
		response.Failed(c, 400, "未成功查询帖子记录", err)
		return
	}
	log.Println("post's userId :", post.UserId, " 你的名字", userId)
	if post.UserId != userId {
		response.Failed(c, 400, "权限不足!", err)
		return
	}
	data, err := json.Marshal(jsons["fileids"])
	if err != nil {
		log.Println("marshal fileids wrong!")
	}
	delete(jsons, "postid")
	delete(jsons, "fileids")
	//更新
	db.Model(&post).Updates(jsons)
	err = db.Model(&post).Update("fileids", string(data)).Error
	if err != nil {
		response.Failed(c, 400, "更新失败", err)
	}
	response.Success(c, 200, "更新成功!", post)
}
