package logic

// User 用户关键数据
type User struct {
	UserType    int    `json:"user_type,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	PassWord    string `json:"password,omitempty"`
	Code        string `json:"code,omitempty"`
}

// UserInfo 用户信息
type UserInfo struct {
	PhoneNumber    string         `json:"phone_number,omitempty"`
	UserBaseInfo   UserBaseInfo   `json:"user_base_info"`
	Habits         Habits         `json:"habits"`
	MedicalHistory MedicalHistory `json:"medical_history"`
	Image          string         `json:"image,omitempty"`     //头像
	NickName       string         `json:"nick_name,omitempty"` //昵称
}

// UserBaseInfo 用户基本信息
type UserBaseInfo struct {
	Name            string  `json:"name,omitempty"`             //姓名
	Age             int     `json:"age,omitempty"`              //年龄
	Gender          bool    `json:"gender,omitempty"`           //性别
	EducationDegree string  `json:"education_degree,omitempty"` //文化程度
	BirthDate       string  `json:"birth_date,omitempty"`       //出生日期 01/02/2006
	Height          float64 `json:"height,omitempty"`           //身高
	Weight          float64 `json:"weight,omitempty"`           //体重
}

// Habits 用户习惯
type Habits struct {
	Smoke        bool `json:"smoke,omitempty"`         //抽烟
	Drink        bool `json:"drink,omitempty"`         //喝酒
	RarelySocial bool `json:"rarely_social,omitempty"` //极少社交
}

// MedicalHistory 既往病史
type MedicalHistory struct {
	Diabetes                       bool   `json:"diabetes,omitempty"`                          //糖尿病
	Apoplexy                       bool   `json:"apoplexy,omitempty"`                          //中风
	Parkinson                      bool   `json:"parkinson,omitempty"`                         //帕金森
	Hypertension                   bool   `json:"hypertension,omitempty"`                      //高血压
	BrainTrauma                    bool   `json:"brain_trauma,omitempty"`                      //脑外伤
	Depression                     bool   `json:"depression,omitempty"`                        //抑郁症
	Hyperlipidemia                 bool   `json:"hyperlipidemia,omitempty"`                    //高血脂
	CardiovascularDisease          bool   `json:"cardiovascular_disease,omitempty"`            //心血管疾病
	CognitiveDisorderFamilyHistory bool   `json:"cognitive_disorder_family_history,omitempty"` //认知症家族史
	Obesity                        bool   `json:"obesity,omitempty"`                           //肥胖症
	Epilepsy                       bool   `json:"epilepsy,omitempty"`                          //癫痫
	Other                          string `json:"other,omitempty"`                             //其他
}
