package dao

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

func SaveCode(phoneNumber, code string) (err error) {
	err = rdb.Set(phoneNumber, code, time.Minute*30).Err()
	return
}

func GetCode(phoneNumber string) (code string, err error) {
	code, err = rdb.Get(phoneNumber).Result()
	return
}

func DelCode(phoneNumber string) (err error) {
	err = rdb.Del(phoneNumber).Err()
	return
}

func (u *User) Register() (err error) {
	fmt.Println(u.PassWord)
	err = db.Create(&u).Error

	var testUser User
	db.Where("phone_number = ?", u.PhoneNumber).First(&testUser)
	fmt.Println(testUser.PassWord)
	return
}

// IsRegister 判断用户是否已存在，true为存在
func IsRegister(phoneNumber string) (IsRegister bool) {
	IsRegister = true
	err := db.Model(User{}).Where("phone_number = ?", phoneNumber).First(User{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			IsRegister = false
		}
	}
	return
}
