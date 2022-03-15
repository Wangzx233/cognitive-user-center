package dao

import "time"

func SavePhoneNumberByRT(RT, phoneNumber string) (err error) {
	err = rdb.Set(RT, phoneNumber, time.Hour*30).Err()
	return
}

func GetPhoneNumberByRT(RT string) (phoneNumber string, err error) {
	phoneNumber, err = rdb.Get(RT).Result()
	return
}
