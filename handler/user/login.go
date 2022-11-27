package handler

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"
	"wechatdemo/utils"
	"wechatdemo/verify"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("0123456789")

func getRandName() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return "user" + string(b)
}
func ReorLo(c *gin.Context) {
	//获取参数
	var verifycodeSend, verifyCode model.VerifyCode
	if err := c.ShouldBind(&verifycodeSend); err != nil {
		log.Println("获取email和code失败")
		response.Failed(c, 400, "获取email和code失败", "")
		return
	}
	//连接数据库
	DB := database.Get()
	var user model.User
	var seconds = 600 //60s
	var t = time.Now()
	email := verifycodeSend.Email
	code := verifycodeSend.Code
	//查询数据库中是否有该用户
	DB.Where("email = ?", email).First(&user)

	if user.ID == 0 { // 说明不存在，进行注册
		DB.Where("email = ?", email).Last(&verifyCode)

		if verifyCode.ID == 0 { //验证码不存在
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "验证码不存在!"})
			return
		}

		if verifyCode.ID != 0 && verifyCode.CreatedAt.Add(time.Duration(seconds)*time.Second).Before(t) { //超时60s以上
			DB.Where("email = ?", email).Delete(&verifyCode)
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "验证码超时!"})
			return
		}

		if verifyCode.ID != 0 && verifyCode.Code == code { //一切正常
			DB.Where("email = ?", email).Delete(&verifyCode)
			user = model.User{
				Email: email,
				Name:  getRandName(),
			}

			DB.Create(&user) //存入数据库
			//c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功!"})
			token, err := verify.ReleaseToken(user) //获取token
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "mes": "生成token失败!"})
				return
			}
			c.Header("Authorization", token)
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功!", "data": token})
		} else {
			DB.Where("email = ?", email).Delete(&verifyCode)
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "验证码无效!"})
			return
		}
	} else { //说明存在,进行登录
		DB.Where("email = ?", email).Last(&verifyCode)

		if verifyCode.ID == 0 { //验证码不存在
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "验证码不存在!"})
			return
		}

		if verifyCode.ID != 0 && verifyCode.CreatedAt.Add(time.Duration(seconds)*time.Second).Before(t) { //超时60s以上
			DB.Where("email = ?", email).Delete(&verifyCode)
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "验证码超时!"})
			return
		}

		if verifyCode.ID != 0 && verifyCode.Code == code { //一切正常
			token, err := verify.ReleaseToken(user) //获取token
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "mes": "生成token失败!"})
				return
			}
			c.Header("Authorization", token)
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功!", "data": token})
		} else {
			DB.Where("email = ?", email).Delete(&verifyCode)
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "验证码无效!"})
			return
		}
	}
}

// 获取验证码
func GetVerifyCode(c *gin.Context) {
	//获取参数
	var user model.VerifyUser
	if err := c.ShouldBind(&user); err != nil {
		response.Failed(c, 400, "参数错误", "")
		return
	}
	code := utils.CreateVerifyCode()

	verifyCode := &model.VerifyCode{
		Email: user.Email,
		Code:  code,
	}

	//连接数据库
	DB := database.Get()

	//存入数据库
	DB.Create(&verifyCode)
	SendEmail(user.Email, code)
	//返回结果
	c.JSON(http.StatusOK, gin.H{"message": "验证码已发出请及时验收"})
}
