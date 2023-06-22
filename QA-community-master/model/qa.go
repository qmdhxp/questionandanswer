package model

import "github.com/jinzhu/gorm"

type Question struct {
	gorm.Model
	Questioner string   `gorm:"column:questioner" form:"questioner"  json:"questioner" binding:"required"`
	Content    string   `gorm:"column:content" form:"content"  json:"content" binding:"required"`
	Title      string   `gorm:"column:title" form:"title" json:"title" binding:"required"`
	Answers    []Answer `gorm:"column:answers" form:"answers"  json:"answers"`
}

type Answer struct {
	gorm.Model
	Answerer   string `gorm:"column:answerer" form:"answerer"  json:"answerer" binding:"required"`
	QuestionID uint   `gorm:"column:questionID" form:"questionID"  json:"questionID" binding:"required"`
	Content    string `gorm:"column:content" form:"content"  json:"content" binding:"required"`
	//Comments   []Comment `gorm:"column:comments"`
}

//type Comment struct {
//	gorm.Model
//	Commenter string `gorm:"column:commenter"`
//	AnswerID  uint   `gorm:"column:answerID" form:"answerID"  json:"answerID" binding:"required"`
//	Content   string `gorm:"column:content" form:"content"  json:"content" binding:"required"`
//}

type ModifyQuestion struct {
	ID      uint   `gorm:"column:id" form:"id" json:"id" binding:"required"`
	Content string `gorm:"column:content" form:"content"  json:"content" binding:"required"`
	Title   string `gorm:"column:title" form:"title" json:"title" binding:"required"`
}

type ModifyAnswer struct {
	ID      uint   `gorm:"column:id" form:"id" json:"id" binding:"required"`
	Content string `gorm:"column:content" form:"content"  json:"content" binding:"required"`
}

type Delete struct {
	ID uint `gorm:"column:id" form:"id" json:"id" binding:"required"`
}
