package post

import (
	"encoding/json"
	"log"
	"wechatdemo/database"
	databaseuser "wechatdemo/database/user"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	db := database.Get()
	userId := c.GetUint("user")
	//获取参数
	var get_post model.GetPost
	//fileids := c.PostFormArray("fileids")
	//err := c.ShouldBind(&fileids)
	//if err != nil {
	//	log.Println("获取数组出错!", err)
	//}
	//log.Println(fileids)
	if err := c.ShouldBind(&get_post); err != nil {
		response.Failed(c, 400, "参数错误", "")
		return
	}
	log.Println(get_post.FileId)
	// avatar := c.PostForm("avatar")
	// title := c.PostForm("title")
	// qq := c.BindJSON("qq")
	// wx := c.PostForm("wx")
	// content := c.PostForm("content")
	// price := c.PostForm("price")
	// location := c.PostForm("location")
	// tag := c.PostForm("tag")
	my_avatar, _, _ := databaseuser.GetUserDetailById(userId)
	var post = model.Post{
		UserId:   get_post.UserId,
		Avatar:   my_avatar,
		Title:    get_post.Title,
		Content:  get_post.Content,
		Price:    get_post.Price,
		Location: get_post.Location,
		Thumb:    get_post.Thumb,
		Reply:    get_post.Reply,
		Follow:   get_post.Follow,
		Tag:      get_post.Tag,
	}
	log.Println("创建帖子的tag为", post.Tag)
	if post.Content == "" || post.Title == "" {
		response.Failed(c, 400, "content或title未给出", nil)
		return
	}
	data, err := json.Marshal(get_post.Fileids)
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
