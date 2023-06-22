package main

import (
	"QA_community/api"
	"QA_community/boot"
)

func main() {
	boot.InitMysql()
	boot.InitRedis()
	api.InitRouter()
}
