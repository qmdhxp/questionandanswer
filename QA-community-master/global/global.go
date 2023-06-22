package global

import "github.com/jinzhu/gorm"

var (
	GlobalDb      *gorm.DB
	SecretSignKey = []byte("cola22") //JWT秘钥
)
