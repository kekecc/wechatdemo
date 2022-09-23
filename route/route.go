package route

import (
	"net/http"
	comments "wechatdemo/handler/comment"
	follows "wechatdemo/handler/follow"
	"wechatdemo/handler/jpg"
	posts "wechatdemo/handler/post"
	users "wechatdemo/handler/user"
	"wechatdemo/middle"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	route := gin.Default()
	route.StaticFS("/static", http.Dir("./static"))
	user := route.Group("/user")
	{
		user.POST("/login", users.ReorLo)
		user.POST("/getcode", users.GetVerifyCode)
		user.GET("/getUserDetail", middle.AuthJWT(), users.GetUserDetail)
		user.PUT("/updateUserDetail", middle.AuthJWT(), users.UpdateUserDetail)
		user.GET("/GetUserPostsSum", middle.AuthJWT(), users.GetUserPostsSum)
		user.GET("/GetUserCommentsSum", middle.AuthJWT(), users.GetUserCommentsSum)
		user.GET("/GetUserFollowsSum", middle.AuthJWT(), users.GetUserFollowsSum)
	}
	post := route.Group("/post")
	{
		post.GET("/getPostsList", posts.GetPostList)
		post.Use(middle.AuthJWT())
		post.POST("/create", posts.Create)
		post.PUT("/update", posts.Update)
		post.DELETE("/delete", posts.Delete)
		post.POST("/thumb", follows.Thumb)
		post.GET("/selectByTitle", posts.GetPostListByTitle)
		//比postlist多一个title字段，get form方式
		post.GET("/selectByTag", posts.GetPostListByTag)
		post.GET("/getPostListByUser", posts.GetPostListByUser)
		//比postlist多一个tag字段，get form方式

		post.GET("/getPostByid", posts.GetPostByid)

	}
	comment := route.Group("/comment", middle.AuthJWT())
	{

		//comment.Use(middle.AuthJWT())
		comment.POST("/create", comments.CreateComment)
		comment.GET("/getCommentListByPost", comments.GetCommentListByPost)
		comment.GET("/getCommentListByUser", comments.GetCommentListByUser)
		comment.GET("/getCommentListByMyPost", comments.GetCommentListByMyself)
		comment.GET("/getReCommentListByUser")
		comment.DELETE("/delete", comments.Delete)
	}
	follow := route.Group("/follow", middle.AuthJWT())
	{
		follow.POST("/follow", follows.Follow)
		follow.DELETE("/deletefollow", follows.DeleteFollow)
		follow.GET("/getFollowList", follows.GetFollowList)
	}
	image := route.Group("/jpg", middle.AuthJWT())
	{
		image.POST("/download", jpg.DownloadJpg)
	}
	route.GET("/comment/getUserByCommentid", comments.GetUserByCommentID)
	return route
}
