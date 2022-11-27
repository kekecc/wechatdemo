package model

import (
	"time"

	"gorm.io/gorm"
)

// 储存的帖子也是创建帖子也是更改帖子的参数
type Post struct {
	gorm.Model
	//UserName string `json:"userName" postform:"userName"` //创建帖子的人
	UserId   uint   `json:"userId" postform:"userId" gorm:"user_id"`
	FileId   string `json:"fileid" postform:"fileid" gorm:"column:fileid"`
	Avatar   string `json:"avatar" postform:"avatar"`
	Title    string `json:"title" postform:"title" binding:"required"`
	Content  string `json:"content" postform:"content" binding:"required"`
	Price    string `json:"price" postform:"price"`
	Location string `json:"location" postform:"location"`
	Thumb    int    `json:"thumb" postform:"thumb"`
	Reply    int    `json:"reply" postform:"reply"`
	Follow   int    `json:"follow" postform:"follow"`
	Tag      string `json:"tag" postform:"tag"`
}

type GetPost struct {
	gorm.Model
	//UserName string `json:"userName" postform:"userName"` //创建帖子的人
	Fileids  []string `json:"fileids" postform:"fileids" `
	UserId   uint     `json:"userId" postform:"userId" gorm:"user_id"`
	FileId   string   `json:"fileid" postform:"fileid" gorm:"column:fileid"`
	Avatar   string   `json:"avatar" postform:"avatar"`
	Title    string   `json:"title" postform:"title" binding:"required"`
	Content  string   `json:"content" postform:"content" binding:"required"`
	Price    string   `json:"price" postform:"price"`
	Location string   `json:"location" postform:"location"`
	Thumb    int      `json:"thumb" postform:"thumb"`
	Reply    int      `json:"reply" postform:"reply"`
	Follow   int      `json:"follow" postform:"follow"`
	Tag      string   `json:"tag" postform:"tag"`
}

// 返回响应的帖子
type ResponsePost struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UserName  string
	Userid    uint
	Fileid    []string
	Avatar    string
	Title     string
	QQ        string
	Wx        string
	Content   string
	Price     string
	Location  string
	Thumb     int
	Reply     int
	Follow    int
	IsThumb   bool `json:"isThumb"`
	IsReplied bool `json:"isReplied"`
	IsFollow  bool `json:"isFollow"`
	Tag       string
}

// 正常询问
type ListType struct {
	Mode   string `uri:"mode" json:"mode" form:"mode" `
	Limit  int    `uri:"limit" json:"limit" form:"limit"`
	Offset int    `uri:"offset" json:"offset" form:"offset" bind:"required"`
}

//点赞

type DeletePost struct {
	Postid uint `json:"postid" postform:"postid" bind:"required"`
}

type Report struct {
	Username string `json:"Username" postform:"Username"`
	Postid   uint   `json:"postid" postform:"postid" bind:"required"`
	Content  string `json:"content" postform:"content" bind:"required"`
}
