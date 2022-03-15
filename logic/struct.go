package logic

// User 用户关键数据
type User struct {
	UserType    int    `json:"user_type,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	PassWord    string `json:"password,omitempty"`
	Code        string `json:"code,omitempty"`
}
