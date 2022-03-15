package logic

import (
	"cognitive-user-center/dao"
	"errors"
	"log"
)

// SaveInfo 新建或更新用户信息
func (userInfo *UserInfo) SaveInfo() (err error) {
	//构建dao层数据
	err = dao.SaveUserInfo(userInfo.ToDaoUserInfo())
	if err != nil {
		log.Println("tx err : ", err)
		err = errors.New("致命错误")
	}

	return
}

func (userInfo *UserInfo) FindInfo() (err error) {
	info, history, habits, baseInfo, err := dao.FindUserInfo(userInfo.PhoneNumber)
	if err != nil {
		log.Println("dao err : ", err)
		err = errors.New("致命错误")
		return
	}

	*userInfo = UserInfo{
		PhoneNumber: userInfo.PhoneNumber,
		UserBaseInfo: UserBaseInfo{
			Name:            baseInfo.Name,
			Age:             baseInfo.Age,
			Gender:          baseInfo.Gender,
			EducationDegree: baseInfo.EducationDegree,
			BirthDate:       baseInfo.BirthDate,
			Height:          baseInfo.Height,
			Weight:          baseInfo.Weight,
		},
		Habits: Habits{
			Smoke:        habits.Smoke,
			Drink:        habits.Drink,
			RarelySocial: habits.RarelySocial,
		},
		MedicalHistory: MedicalHistory{
			Diabetes:                       history.Diabetes,
			Apoplexy:                       history.Apoplexy,
			Parkinson:                      history.Parkinson,
			Hypertension:                   history.Hypertension,
			BrainTrauma:                    history.BrainTrauma,
			Depression:                     history.Depression,
			Hyperlipidemia:                 history.Hyperlipidemia,
			CardiovascularDisease:          history.CardiovascularDisease,
			CognitiveDisorderFamilyHistory: history.CognitiveDisorderFamilyHistory,
			Obesity:                        history.Obesity,
			Epilepsy:                       history.Epilepsy,
			Other:                          history.Other,
		},
		Image:    info.Image,
		NickName: info.NickName,
	}
	return
}
