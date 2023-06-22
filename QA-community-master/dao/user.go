package dao

import (
	"QA_community/global"
	"QA_community/model"
	"strings"
)

// 验证密码格式(至少含有大小写字母，数字，特殊字符中的2种)
func JudgePassword(password string) bool {
	count := 0
	for _, number := range password {
		if number >= 48 && number <= 57 {
			count += 1
			break
		}
	}
	for _, lowerletter := range password {
		if lowerletter >= 97 && lowerletter <= 122 {
			count += 1
			break
		}
	}
	for _, higherletter := range password {
		if higherletter >= 65 && higherletter <= 90 {
			count += 1
			break
		}
	}
	for _, schar := range password {
		if schar < 48 {
			count += 1
			break
		}
		if schar > 57 && schar < 65 {
			count += 1
			break
		}
		if schar > 90 && schar < 97 {
			count += 1
			break
		}
		if schar > 122 {
			count += 1
			break
		}
	}
	if count < 2 {
		return false
	}
	return true
}

// 验证手机号格式
func JudgePhoneNumber(phonenumber string) bool {
	if len([]rune(phonenumber)) != 11 {
		return false
	}
	if strings.HasPrefix(phonenumber, "1") != true {
		return false
	}
	return true
}

// 验证邮箱格式
func JudgeEmail(email string) bool {
	flag := 0
	if strings.HasSuffix(email, "@qq.com") {
		flag = 1
	}
	if strings.HasSuffix(email, "@QQ.com") {
		flag = 1
	}
	if flag == 0 {
		return false
	}
	return true
}

// 检查用户名是否重复
// 若没有这个用户返回 ture，反之返回 false
func CheckUserName(username string) bool {
	var user model.User
	global.GlobalDb.Model(&model.User{}).Where("username=?", username).Find(&user)
	if user.UserName == "" {
		return true
	}
	return false
}

// 检查手机号是否重复
// 若没有这个手机号返回 ture，反之返回 false
func CheckPhoneNumber(phonenumber string) bool {
	var user model.User
	global.GlobalDb.Model(&model.User{}).Where("phonenumber=?", phonenumber).Find(&user)
	if user.PhoneNumber == "" {
		return true
	}
	return false
}

// 检查邮箱是否重复
// 若没有这个邮箱返回 ture，反之返回 false
func CheckEmail(email string) bool {
	var user model.User
	global.GlobalDb.Model(&model.User{}).Where("email=?", email).Find(&user)
	if user.Email == "" {
		return true
	}
	return false
}

// 密码验证(未完善)
func PasswordAuth(username string, password string) bool {
	var user model.User
	global.GlobalDb.Model(&model.User{}).Where("username=?", username).Find(&user)
	if user.Password == password {
		return true
	}
	return false
}

func GetUsernameFromEmail(email string) string {
	var user model.User
	global.GlobalDb.Model(&model.User{}).Where("email=?", email).Find(&user)
	return user.UserName
}

func GetUserFromPhoneNumber(phonenumber string) string {
	var user model.User
	global.GlobalDb.Model(&model.User{}).Where("phonenumber=?", phonenumber).Find(&user)
	return user.UserName
}
