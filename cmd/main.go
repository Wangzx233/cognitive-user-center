package main

import (
	"cognitive-user-center/dao"
	"cognitive-user-center/router"
)

func main() {
	dao.InitMysql()
	dao.InitRedis()
	router.InitRouter()
}
