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

func (userInfo *UserInfo) ToDaoUserInfo() (info dao.UserInfo, history dao.MedicalHistory, habits dao.Habits, baseInfo dao.UserBaseInfo) {
	info = dao.UserInfo{
		PhoneNumber: userInfo.PhoneNumber,
		Image:       userInfo.Image,
		NickName:    userInfo.NickName,
	}

	history = dao.MedicalHistory{
		PhoneNumber:                    userInfo.PhoneNumber,
		Diabetes:                       userInfo.MedicalHistory.Diabetes,
		Apoplexy:                       userInfo.MedicalHistory.Apoplexy,
		Parkinson:                      userInfo.MedicalHistory.Parkinson,
		Hypertension:                   userInfo.MedicalHistory.Hypertension,
		BrainTrauma:                    userInfo.MedicalHistory.BrainTrauma,
		Depression:                     userInfo.MedicalHistory.Depression,
		Hyperlipidemia:                 userInfo.MedicalHistory.Hyperlipidemia,
		CardiovascularDisease:          userInfo.MedicalHistory.CardiovascularDisease,
		CognitiveDisorderFamilyHistory: userInfo.MedicalHistory.CognitiveDisorderFamilyHistory,
		Obesity:                        userInfo.MedicalHistory.Obesity,
		Epilepsy:                       userInfo.MedicalHistory.Epilepsy,
		Other:                          userInfo.MedicalHistory.Other,
	}

	habits = dao.Habits{
		PhoneNumber:  userInfo.PhoneNumber,
		Smoke:        userInfo.Habits.Smoke,
		Drink:        userInfo.Habits.Drink,
		RarelySocial: userInfo.Habits.RarelySocial,
	}

	baseInfo = dao.UserBaseInfo{
		PhoneNumber:     userInfo.PhoneNumber,
		Name:            userInfo.UserBaseInfo.Name,
		Age:             userInfo.UserBaseInfo.Age,
		Gender:          userInfo.UserBaseInfo.Gender,
		EducationDegree: userInfo.UserBaseInfo.EducationDegree,
		BirthDate:       userInfo.UserBaseInfo.BirthDate,
		Height:          userInfo.UserBaseInfo.Height,
		Weight:          userInfo.UserBaseInfo.Weight,
	}
	return
}
