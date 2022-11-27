package post

import (
	"encoding/json"
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	db := database.Get()
	userId := c.GetUint("user")
	//获取参数
	var post model.Post
	var fileids []string
	err := c.ShouldBindJSON(&fileids)
	if err != nil {
		log.Println("获取数组出错!", err)
	}
	log.Println(fileids)
	if err := c.ShouldBind(&post); err != nil {
		response.Failed(c, 400, "参数错误", "")
		return
	}
	// avatar := c.PostForm("avatar")
	// title := c.PostForm("title")
	// qq := c.BindJSON("qq")
	// wx := c.PostForm("wx")
	// content := c.PostForm("content")
	// price := c.PostForm("price")
	// location := c.PostForm("location")
	// tag := c.PostForm("tag")
	log.Println("创建帖子的tag为", post.Tag)
	if post.Content == "" || post.Title == "" {
		response.Failed(c, 400, "content或title未给出", nil)
		return
	}
	data, err := json.Marshal(fileids)
	if err != nil {
		response.Failed(c, 400, "转json失败", nil)
		return
	}
	log.Println(string(data))
	if len(data) != 0 {
		post.FileId = string(data)
	} else {
		post.FileId = ""
	}
	post.UserId = userId
	err = db.Table("post").Create(&post).Error
	if err != nil {
		log.Println("创建帖子失败")
	}
	log.Println("创建帖子", post)
	response.Success(c, 200, "创建帖子成功", post)
}
