package logic

import (
	"cognitive-user-center/dao"
	"errors"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"time"
)

const AccessTokenExpireDuration = time.Hour * 2

var MySecret = []byte("mozzarella")

type MyClaims struct {
	PhoneNumber string `json:"phone_number"`
	jwt.StandardClaims
}

// GenAccessToken 生成JWT
func GenAccessToken(phoneNumber string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		phoneNumber, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(), // 过期时间
			Issuer:    "cognitive-user-center",                          // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// GenRefreshToken 生成JWT
func GenRefreshToken(phoneNumber string) (RT string, err error) {
	RT = uuid.NewV4().String()
	err = dao.SavePhoneNumberByRT(RT, phoneNumber)

	return
}

// ParseAccessToken 解析JWT
func ParseAccessToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
