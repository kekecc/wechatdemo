package post

import (
	"encoding/json"
	"log"
	"wechatdemo/database"
	databaseuser "wechatdemo/database/user"
	"wechatdemo/model"
	"wechatdemo/response"
	"wechatdemo/utils"

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

	// upload
	// forms, err := c.MultipartForm()
	// if err != nil {
	// 	log.Println(err)
	// 	response.Failed(c, 400, "获取所有图片出错", nil)
	// 	return
	// }

	// files := forms.File["files"] // images
	// utils.UploadMutipy("posts", files)

	// cos := utils.GetCos()
	// for i, file := range files {
	// 	url := cos.Object.GetObjectURL("posts" + file.Filename)
	// 	get_post.Fileids[i] = url.String()
	// }
	get_post.Fileids = make([]string, 0)
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

func UploadImage(c *gin.Context) {
	id := c.PostForm("postid")
	if id == "" {
		log.Println("上传图片但是帖子id为空")
		response.Failed(c, 400, "未获取到postid", nil)
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("未获取到file")
		response.Failed(c, 400, "未获取到file", err)
		return
	}

	userId := c.GetUint("user")
	db := database.Get()

	var post model.Post
	if err := db.Model(&model.Post{}).Where("id = ?", id).First(&post).Error; err != nil {
		log.Println("获取post出错", err)
		response.Failed(c, 400, "获取帖子出错", err)
		return
	}

	// exist
	if post.UserId != userId {
		log.Println("用户和帖子不对应")
		response.Failed(c, 400, "用户帖子不对应", nil)
		return
	}

	var ids []string
	if post.FileId != "" {
		err := json.Unmarshal([]byte(post.FileId), &ids)
		if err != nil {
			log.Println("解析json出错!", err)
		}
	} else {
		ids = append(ids, "")
	}

	f, _ := file.Open()
	if err = utils.Upload("post", file.Filename, f); err != nil {
		log.Println(err)
		response.Failed(c, 400, "上传头像错误", nil)
		return
	}
	f.Close()

	cos := utils.GetCos()
	_url := cos.Object.GetObjectURL("post" + file.Filename)

	ids = append(ids, _url.String())

	data, err := json.Marshal(ids)
	if err != nil {
		response.Failed(c, 400, "转json失败", nil)
		return
	}

	db.Model(&model.Post{}).Where("id = ?", id).Update("fileid", string(data))
	response.Success(c, 200, "上传图片成功", nil)
}
