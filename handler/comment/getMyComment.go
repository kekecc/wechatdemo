package comment

import (
	databasecomment "wechatdemo/database/comment"
	databaseuser "wechatdemo/database/user"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetCommentListByMyself(c *gin.Context) {
	userid := c.GetUint("user")
	offset := com.StrTo(c.Query("offset")).MustInt()
	limit := com.StrTo(c.Query("limit")).MustInt()
	comments := databasecomment.GetCommentByMyself(c, userid, limit, offset)
	if comments == nil {
		response.Failed(c, 400, "获取回复失败!", nil)
	}
	var replyComments []model.ReplyComments
	for _, comment := range comments {
		//log.Println("获取第", i+1, "个评论")
		if comment.UserId == userid {
			continue
		}
		var replyComment model.ReplyComments
		replyComment.ID = comment.Postid
		replyComment.Content = comment.Content
		userName, err := databaseuser.GetUserNameByID(comment.UserId)
		if err == nil {
			replyComment.UserName = userName
		} else {
			continue
		}
		replyComment.Fileid, replyComment.QQ, replyComment.Wx = databaseuser.GetUserDetailById(comment.UserId)
		replyComments = append(replyComments, replyComment)
	}
	response.Success(c, 200, "成功", replyComments)
}
