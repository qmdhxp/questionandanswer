package api

import (
	"QA_community/tools/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/api/register", UserRegister)             // 注册
	r.POST("/api/namelogin", UserNameLogin)           // 使用账号密码登录
	r.POST("/api/emaillogin", EmailLogin)             // 使用邮箱密码登录
	r.POST("/api/phonenumberlogin", PhoneNumberLogin) // 使用手机号密码登录
	QARouter := r.Group("/api/qa")                    // 问答分组(q为问题,a为回答)
	{
		QARouter.Use(middleware.AuthMiddleWare())
		QARouter.POST("/createq", CreateQuestion)   // 发布问题
		QARouter.POST("/createa", CreateAnswer)     // 发布回答
		QARouter.GET("/viewq", ViewOwnQuestions)    // 查看自己发布的问题及回答
		QARouter.GET("/viewa", ViewOwnAnswers)      // 查看自己发布的回答
		QARouter.PUT("/modifyq", ModifyQuestion)    // 修改问题
		QARouter.PUT("/modifya", ModifyAnswer)      // 修改回答
		QARouter.DELETE("/deleteq", DeleteQuestion) // 删除问题
		QARouter.DELETE("/deletea", DeleteAnswer)   // 删除回答
	}
	r.Run(":8090") // 跑在 8090 端口上
}
