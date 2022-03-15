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
}
