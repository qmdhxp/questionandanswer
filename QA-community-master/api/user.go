package api

import (
	"QA_community/dao"
	"QA_community/global"
	"QA_community/model"
	"QA_community/tools/checkresponse"
	"QA_community/tools/jwttoken"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	//获取registerform表单数据
	var registerform model.UserRegister
	err := c.ShouldBindJSON(&registerform)
	if err != nil {
		checkresponse.ResFail(c, "参数绑定失败")
		return
	}
	var flag bool
	//验证密码格式
	flag = dao.JudgePassword(registerform.Password)
	if !flag {
		checkresponse.ResFail(c, "密码格式错误")
		return
	}
	//验证手机号格式
	flag = dao.JudgePhoneNumber(registerform.PhoneNumber)
	if !flag {
		checkresponse.ResFail(c, "手机号格式错误")
		return
	}
	//验证邮箱格式
	flag = dao.JudgeEmail(registerform.Email)
	if !flag {
		checkresponse.ResFail(c, "邮箱格式有误")
		return
	}
	//检查用户名是否重复
	flag = dao.CheckUserName(registerform.UserName)
	if !flag {
		checkresponse.ResFail(c, "用户名已存在")
		return
	}
	//检查手机号是否重复
	flag = dao.CheckPhoneNumber(registerform.PhoneNumber)
	if !flag {
		checkresponse.ResFail(c, "手机号已存在")
		return
	}
	//检查邮箱是否重复
	flag = dao.CheckEmail(registerform.Email)
	if !flag {
		checkresponse.ResFail(c, "手机号已存在")
		return
	}
	//存入数据库
	var user model.User
	user = model.User{
		UserName:    registerform.UserName,
		Password:    registerform.Password,
		Email:       registerform.Email,
		PhoneNumber: registerform.PhoneNumber,
	}
	global.GlobalDb.Model(&model.User{}).Create(&user)
	checkresponse.ResSuccess(c, "注册成功")
}

func PhoneNumberLogin(c *gin.Context) {
	var loginform model.PhoneLogin
	err := c.ShouldBindJSON(&loginform)
	if err != nil {
		checkresponse.ResFail(c, "参数绑定失败")
		return
	}
	//检查手机号是否存在
	flag := dao.CheckPhoneNumber(loginform.PhoneNumber)
	if flag {
		checkresponse.ResFail(c, "手机号不存在")
		return
	}
	//密码验证
	flag = dao.PasswordAuth(dao.GetUserFromPhoneNumber(loginform.PhoneNumber), loginform.Password)
	if !flag {
		checkresponse.ResFail(c, "密码错误")
		return
	}
	username := dao.GetUserFromPhoneNumber(loginform.PhoneNumber)
	//生成JWTtoken
	tokenString, _ := jwttoken.CreateToken(username)
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"data":     gin.H{"token": tokenString},
		"username": username,
	})
}

func EmailLogin(c *gin.Context) {
	var loginform model.EmailLogin
	err := c.ShouldBindJSON(&loginform)
	if err != nil {
		checkresponse.ResFail(c, "参数绑定失败")
		return
	}
	//检查邮箱是否存在
	flag := dao.CheckEmail(loginform.Email)
	if flag {
		checkresponse.ResFail(c, "邮箱不存在")
		return
	}
	//密码验证
	flag = dao.PasswordAuth(dao.GetUsernameFromEmail(loginform.Email), loginform.Password)
	if !flag {
		checkresponse.ResFail(c, "密码错误")
		return
	}
	username := dao.GetUsernameFromEmail(loginform.Email)
	//生成JWTtoken
	tokenString, _ := jwttoken.CreateToken(username)
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"data":     gin.H{"token": tokenString},
		"username": username,
	})
}

func UserNameLogin(c *gin.Context) {
	var loginform model.UserNameLogin
	err := c.ShouldBindJSON(&loginform)
	if err != nil {
		checkresponse.ResFail(c, "参数绑定失败")
		return
	}
	//检查用户名是否存在
	flag := dao.CheckUserName(loginform.UserName)
	if flag {
		checkresponse.ResFail(c, "用户名不存在")
		return
	}
	//密码验证
	flag = dao.PasswordAuth(loginform.UserName, loginform.Password)
	if !flag {
		checkresponse.ResFail(c, "密码错误")
		return
	}
	username := loginform.UserName
	//生成JWTtoken
	tokenString, _ := jwttoken.CreateToken(username)
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"data":     gin.H{"token": tokenString},
		"username": username,
	})
}
