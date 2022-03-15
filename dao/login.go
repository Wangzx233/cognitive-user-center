package dao

func (u *User) LoginByPassword() (err error) {
	err = db.Model(User{}).Where("phone_number = ? and pass_word = ?", u.PhoneNumber, u.PassWord).First(&u).Error
	return
}

func FindPasswordByPhoneNumber(phoneNumber string) (password []byte, err error) {
	var tempUser User
	err = db.Where("phone_number = ?", phoneNumber).First(&tempUser).Error
	password = tempUser.PassWord

	return
}