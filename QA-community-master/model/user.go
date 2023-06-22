package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName    string `gorm:"column:username" form:"username"  json:"username" binding:"required"`
	Email       string `gorm:"column:email" form:"email" json:"email" binding:"required"`
	Password    string `gorm:"column:password" form:"password" json:"password" binding:"required"`
	PhoneNumber string `gorm:"column:phonenumber" form:"phonenumber" json:"phonenumber" binding:"required"`
}

type UserRegister struct {
	UserName    string `gorm:"column:username" form:"username"  json:"username" binding:"required"`
	Password    string `gorm:"column:password" form:"password" json:"password" binding:"required"`
	Email       string `gorm:"column:email" form:"email" json:"email" binding:"required"`
	PhoneNumber string `gorm:"column:phonenumber" form:"phonenumber" json:"phonenumber" binding:"required"`
}

type PhoneLogin struct {
	PhoneNumber string `gorm:"column:phonenumber" form:"phonenumber" json:"phonenumber" binding:"required"`
	Password    string `gorm:"column:password" form:"password" json:"password" binding:"required"`
}

type UserNameLogin struct {
	UserName string `gorm:"column:username" form:"username"  json:"username" binding:"required"`
	Password string `gorm:"column:password" form:"password" json:"password" binding:"required"`
}

type EmailLogin struct {
	Email    string `gorm:"column:email" form:"email" json:"email" binding:"required"`
	Password string `gorm:"column:password" form:"password" json:"password" binding:"required"`
}
