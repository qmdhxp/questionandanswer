package api

import (
	"QA_community/dao"
	"QA_community/global"
	"QA_community/model"
	"QA_community/tools/checkresponse"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateQuestion(c *gin.Context) {
	//接收前端发来的数据
	var question model.Question
	err := c.ShouldBindJSON(&question)
	if err != nil {
		checkresponse.ResFail(c, "参数绑定失败")
		return
	}
	//通过c.Get方法获取的值的类型是interface{}，因此需要根据实际情况进行类型断言或转换，以便正确使用该值。
	usernameRaw, exists := c.Get("username")
	if !exists {
		checkresponse.ResFail(c, "无法获取用户名")
		return
	}
	username, ok := usernameRaw.(string)
	if !ok {
		checkresponse.ResFail(c, "用户名类型断言失败")
		return
	}
	fmt.Println(username)
	//存入数据库
	depositquestion := model.Question{
		Questioner: username,
		Title:      question.Title,
		Content:    question.Content,
	}
	global.GlobalDb.Model(&model.Question{}).Create(&depositquestion)
	checkresponse.ResSuccess(c, "你已成功提问")
}

func CreateAnswer(c *gin.Context) {
	//接收前端发来的数据
	var answer model.Answer
	err := c.ShouldBindJSON(&answer)
	if err != nil {
		checkresponse.ResFail(c, "参数绑定失败")
		return
	}
	//通过c.Get方法获取的值的类型是interface{}，因此需要根据实际情况进行类型断言或转换，以便正确使用该值。
	usernameRaw, exists := c.Get("username")
	if !exists {
		checkresponse.ResFail(c, "无法获取用户名")
		return
	}
	username, ok := usernameRaw.(string)
	if !ok {
		checkresponse.ResFail(c, "用户名类型断言失败")
		return
	}
	//存入数据库
	depositanswer := model.Answer{
		Answerer:   username,
		QuestionID: answer.QuestionID,
		Content:    answer.Content,
	}
	global.GlobalDb.Model(&model.Question{}).Create(&depositanswer)
	checkresponse.ResSuccess(c, "回答成功")
}

//	func CreateComment(c *gin.Context) {
//		//接收前端发来的数据
//		var comment model.Comment
//		err := c.ShouldBindJSON(&comment)
//		if err != nil {
//			checkresponse.ResFail(c, "参数绑定失败")
//			return
//		}
//		//通过c.Get方法获取的值的类型是interface{}，因此需要根据实际情况进行类型断言或转换，以便正确使用该值。
//		usernameRaw, exists := c.Get("username")
//
//		if !exists {
//			checkresponse.ResFail(c, "无法获取用户名")
//			return
//		}
//
//		username, ok := usernameRaw.(string)
//
//		if !ok {
//			checkresponse.ResFail(c, "用户名类型断言失败")
//			return
//		}
//
//		//存入数据库
//		depositcomment := model.Comment{
//			Commenter: username,
//			AnswerID:  comment.AnswerID,
//			Content:   comment.Content,
//		}
//		global.GlobalDb.Model(&model.Question{}).Create(&depositcomment)
//	}
func ViewOwnQuestions(c *gin.Context) {
	usernameRaw, exists := c.Get("username")
	if !exists {
		checkresponse.ResFail(c, "无法获取用户名")
		return
	}
	username, ok := usernameRaw.(string)
	if !ok {
		checkresponse.ResFail(c, "用户名类型断言失败")
		return
	}
	questions := dao.FindQuestions(username)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": questions,
	})
}

func ViewOwnAnswers(c *gin.Context) {
	usernameRaw, exists := c.Get("username")
	if !exists {
		checkresponse.ResFail(c, "无法获取用户名")
		return
	}
	username, ok := usernameRaw.(string)
	if !ok {
		checkresponse.ResFail(c, "用户名类型断言失败")
		return
	}
	answers := dao.FindAnswers(username)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": answers,
	})
}

func ModifyQuestion(c *gin.Context) {
	var modifyquestion model.ModifyQuestion
	err := c.ShouldBindJSON(&modifyquestion)
	if err != nil {
		checkresponse.ResFail(c, "参数绑定失败")
		return
	}
	//进行身份验证
	usernameRaw, exists := c.Get("username")
	if !exists {
		checkresponse.ResFail(c, "无法获取用户名")
		return
	}
	username, ok := usernameRaw.(string)
	if !ok {
		checkresponse.ResFail(c, "用户名类型断言失败")
		return
	}
	if !dao.QAuthIdentity(username, modifyquestion.ID) {
		checkresponse.ResFail(c, "修改失败,这个问题不属于你")
		return
	}
	dao.QModify(modifyquestion.Content, modifyquestion.ID)
	checkresponse.ResSuccess(c, "问题修改成功")
}

func ModifyAnswer(c *gin.Context) {
	var modifyanswer model.ModifyAnswer
	err := c.ShouldBindJSON(&modifyanswer)
	if err != nil {
		checkresponse.ResFail(c, "参数绑定失败")
		return
	}
	//进行身份验证
	usernameRaw, exists := c.Get("username")
	if !exists {
		checkresponse.ResFail(c, "无法获取用户名")
		return
	}
	username, ok := usernameRaw.(string)
	if !ok {
		checkresponse.ResFail(c, "用户名类型断言失败")
		return
	}
	if !dao.AAuthIdentity(username, modifyanswer.ID) {
		checkresponse.ResFail(c, "修改失败,这个回答不属于你")
		return
	}
	dao.AModify(modifyanswer.Content, modifyanswer.ID)
	checkresponse.ResSuccess(c, "回答修改成功")
}

func DeleteQuestion(c *gin.Context) {
	var deleteq model.Delete
	err := c.ShouldBindJSON(&deleteq)
	if err != nil {
		checkresponse.ResFail(c, "参数绑定失败")
		return
	}
	//进行身份验证
	usernameRaw, exists := c.Get("username")
	if !exists {
		checkresponse.ResFail(c, "无法获取用户名")
		return
	}
	username, ok := usernameRaw.(string)
	if !ok {
		checkresponse.ResFail(c, "用户名类型断言失败")
		return
	}
	if !dao.QAuthIdentity(username, deleteq.ID) {
		checkresponse.ResFail(c, "删除失败,这个问题不属于你")
		return
	}
	dao.QDelete(deleteq.ID)
	checkresponse.ResSuccess(c, "删除成功")
}

func DeleteAnswer(c *gin.Context) {
	var deletea model.Delete
	err := c.ShouldBindJSON(&deletea)
	if err != nil {
		checkresponse.ResFail(c, "参数绑定失败")
		return
	}
	//进行身份验证
	usernameRaw, exists := c.Get("username")
	if !exists {
		checkresponse.ResFail(c, "无法获取用户名")
		return
	}
	username, ok := usernameRaw.(string)
	if !ok {
		checkresponse.ResFail(c, "用户名类型断言失败")
		return
	}
	if !dao.AAuthIdentity(username, deletea.ID) {
		checkresponse.ResFail(c, "删除失败,这个回答不属于你")
		return
	}
	dao.ADelete(deletea.ID)
	checkresponse.ResSuccess(c, "删除成功")
}
