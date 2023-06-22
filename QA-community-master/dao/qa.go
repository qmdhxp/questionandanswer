package dao

import (
	"QA_community/global"
	"QA_community/model"
)

func FindQuestions(username string) []model.Question {
	var questions []model.Question
	//查询方法一
	global.GlobalDb.Model(&model.Question{}).Preload("Answers").Where("questioner=?", username).Find(&questions)
	//查询方法二
	//global.GlobalDb.Preload("Answers", func(d *gorm.DB) *gorm.DB {
	//	return d.Where("questioner=?", username)
	//}).Find(&questions)
	return questions
}

func FindAnswers(username string) []model.Answer {
	var answers []model.Answer
	global.GlobalDb.Model(&model.Answer{}).Where("Answerer = ?", username).Find(&answers)
	return answers
}

func QAuthIdentity(username string, id uint) bool {
	var question model.Question
	global.GlobalDb.Model(&model.Question{}).Where("questioner = ? AND id = ?", username, id).First(&question)
	if question.Questioner == username {
		return true
	}
	return false
}

func AAuthIdentity(username string, id uint) bool {
	var answer model.Answer
	global.GlobalDb.Model(&model.Answer{}).Where("answerer = ? AND id = ?", username, id).First(&answer)
	if answer.Answerer == username {
		return true
	}
	return false
}

func QModify(content string, id uint) {
	global.GlobalDb.Model(&model.Question{}).Where("id = ?", id).Update("content", content)

}

func AModify(content string, id uint) {
	global.GlobalDb.Model(&model.Answer{}).Where("id = ?", id).Update("content", content)
}

func QDelete(id uint) {
	var question model.Question
	var answers []model.Answer
	global.GlobalDb.Model(&model.Answer{}).Unscoped().Where("questionID=?", id).Delete(&answers)
	global.GlobalDb.Model(&model.Question{}).Unscoped().Where("id=?", id).Delete(&question)
}

func ADelete(id uint) {
	var answer model.Answer
	global.GlobalDb.Model(&model.Answer{}).Unscoped().Where("id=?", id).Delete(&answer)
}
