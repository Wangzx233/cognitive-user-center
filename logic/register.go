package logic

import (
	"cognitive-user-center/dao"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func (u *User) Register() (err error) {
	code, err := dao.GetCode(u.PhoneNumber)
	if err != nil {
		log.Println("get code err : ", err)
		err = errors.New("验证码错误")
		return
	}

	if code != u.Code {
		err = errors.New("验证码错误")
		return
	}

	daoUser := u.ToDaoUser()

	err = daoUser.Register()
	if err != nil {
		log.Println("dao err : ", err)
		err = errors.New("致命错误")
		return
	}

	return
}

func SendRegisterCode(phoneNumber string) (err error) {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(899999) + 100000

	code := strconv.Itoa(number)

	//todo:发送短信
	log.Println(code)

	err = dao.SaveCode(phoneNumber, code)
	if err != nil {
		log.Println("save code err : ", err)
		err = errors.New("sorry")
	}

	return
}

func IsRegister(phoneNumber string) (IsRegister bool) {
	return dao.IsRegister(phoneNumber)
}
