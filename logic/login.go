package logic

import (
	"cognitive-user-center/dao"
	"cognitive-user-center/rsa"
	"errors"
	"log"
)

// LoginByCode 验证码登陆
func (u *User) LoginByCode() (err error) {
	code, err := dao.GetCode(u.PhoneNumber)
	if err != nil {
		log.Println("get code err : ", err)
		err = errors.New("验证码错误")
		return
	}

	if code != u.Code {
		err = errors.New("验证码错误")
	}

	return
}

// LoginByPassword 密码登陆
func (u *User) LoginByPassword() (err error) {

	password, err := dao.FindPasswordByPhoneNumber(u.PhoneNumber)
	if err != nil {
		log.Println("get password err : ", err)
		err = errors.New("err")
		return
	}

	p := rsa.Decrypt(password, "private.pem")

	if u.PassWord != string(p) {
		err = errors.New("密码错误")
	}

	return
}
