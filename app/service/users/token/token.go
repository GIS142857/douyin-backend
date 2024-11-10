package token

import (
	"douyin-backend/app/global/my_errors"
	"douyin-backend/app/global/variable"
	"douyin-backend/app/http/middleware/my_jwt"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// CreateUserFactory 创建 userToken 工厂
func CreateUserFactory() *userToken {
	return &userToken{
		userJwt: my_jwt.CreateMyJWT(variable.ConfigYml.GetString("Token.JwtTokenSignKey")),
	}
}

type userToken struct {
	userJwt *my_jwt.JwtSign
}

// GenerateToken 生成token
func (u *userToken) GenerateToken(uid int64, nickname string, phone, password string, expireAt int64) (tokens string, err error) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	customClaims := my_jwt.CustomClaims{
		UID:      uid,
		NickName: nickname,
		Phone:    phone,
		Password: password,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,       // 生效开始时间
			ExpiresAt: time.Now().Unix() + expireAt, // 失效截止时间
		},
	}
	return u.userJwt.CreateToken(customClaims)
}

// ParseToken 将 token 解析为绑定时传递的参数
func (u *userToken) ParseToken(tokenStr string) (CustomClaims my_jwt.CustomClaims, err error) {
	if customClaims, err := u.userJwt.ParseToken(tokenStr); err == nil {
		return *customClaims, nil
	} else {
		return my_jwt.CustomClaims{}, errors.New(my_errors.ErrorsParseTokenFail)
	}
}
