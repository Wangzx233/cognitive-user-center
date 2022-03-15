package dao

// User 用户关键数据
type User struct {
	UserType    int
	PhoneNumber string `gorm:"primaryKey"`
	PassWord    []byte
}

// UserBaseInfo 用户基本信息
type UserBaseInfo struct {
	PhoneNumber     string  `gorm:"primaryKey"`
	Name            string  //姓名
	Age             int     //年龄
	Gender          bool    //性别
	EducationDegree string  //文化程度
	BirthDate       string  //出生日期 01/02/2006
	Height          float64 //身高
	Weight          float64 //体重
}

// UserInfo 用户信息
type UserInfo struct {
	PhoneNumber string `gorm:"primaryKey"`
	Image       string //头像
	NickName    string //昵称
}

// Habits 用户习惯
type Habits struct {
	PhoneNumber  string `gorm:"primaryKey"`
	Smoke        bool   //抽烟
	Drink        bool   //喝酒
	RarelySocial bool   //极少社交
}

// MedicalHistory 既往病史
type MedicalHistory struct {
	PhoneNumber                    string `gorm:"primaryKey"`
	Diabetes                       bool   //糖尿病
	Apoplexy                       bool   //中风
	Parkinson                      bool   //帕金森
	Hypertension                   bool   //高血压
	BrainTrauma                    bool   //脑外伤
	Depression                     bool   //抑郁症
	Hyperlipidemia                 bool   //高血脂
	CardiovascularDisease          bool   //心血管疾病
	CognitiveDisorderFamilyHistory bool   //认知症家族史
	Obesity                        bool   //肥胖症
	Epilepsy                       bool   //癫痫
	Other                          string //其他
}
