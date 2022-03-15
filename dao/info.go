package dao

import "gorm.io/gorm"

// FindUserInfo 查找用户信息
func FindUserInfo(phoneNumber string) (info UserInfo, history MedicalHistory, habits Habits, baseInfo UserBaseInfo, err error) {
	if err = db.Where("phone_number = ?", phoneNumber).First(&info).Error; err != nil {
		return
	}

	if err = db.Where("phone_number = ?", phoneNumber).First(&history).Error; err != nil {
		return
	}

	if err = db.Where("phone_number = ?", phoneNumber).First(&habits).Error; err != nil {
		return
	}

	if err = db.Where("phone_number = ?", phoneNumber).First(&baseInfo).Error; err != nil {
		return
	}
	return
}

// SaveUserInfo 保存用户信息
func SaveUserInfo(info UserInfo, history MedicalHistory, habits Habits, baseInfo UserBaseInfo) (err error) {
	err = db.Model(UserInfo{}).Where("phone_number = ?", info.PhoneNumber).First(User{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = CreateUserInfo(info, history, habits, baseInfo)
		}
		return
	}
	err = UpdateUserInfo(info, history, habits, baseInfo)
	return
}

// CreateUserInfo 创建用户信息
func CreateUserInfo(info UserInfo, history MedicalHistory, habits Habits, baseInfo UserBaseInfo) (err error) {
	err = db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&info).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&history).Error; err != nil {
			return err
		}

		if err := tx.Create(&habits).Error; err != nil {
			return err
		}
		if err := tx.Create(&baseInfo).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
	return
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(info UserInfo, history MedicalHistory, habits Habits, baseInfo UserBaseInfo) (err error) {
	err = db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Save(&info).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Save(&history).Error; err != nil {
			return err
		}

		if err := tx.Save(&habits).Error; err != nil {
			return err
		}
		if err := tx.Save(&baseInfo).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
	return
}
