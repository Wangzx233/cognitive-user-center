package logic

import (
	"cognitive-user-center/dao"
	"cognitive-user-center/rsa"
)

// ToDaoUser logic层user转dao层user并把密码加密
func (u *User) ToDaoUser() (daoUser dao.User) {
	//phoneNumberSrc := []byte(u.PhoneNumber)
	//encryptPhoneNumber := rsa.Encrypt(phoneNumberSrc, "public.pem")

	passwordSrc := []byte(u.PassWord)
	encryptPassword := rsa.Encrypt(passwordSrc, "public.pem")

	daoUser = dao.User{
		UserType:    u.UserType,
		PhoneNumber: u.PhoneNumber,
		PassWord:    encryptPassword,
	}

	return
}
