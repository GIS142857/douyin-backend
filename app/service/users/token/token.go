package token

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/global/my_errors"
	"douyin-backend/app/global/variable"
	"douyin-backend/app/http/middleware/my_jwt"
	"douyin-backend/app/model/user"
	"douyin-backend/app/service/users/token_cache_redis"
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
func (u *userToken) GenerateToken(uid int64, nickname string, phone string, expireAt int64) (tokens string, err error) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	customClaims := my_jwt.CustomClaims{
		UID:      uid,
		NickName: nickname,
		Phone:    phone,
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

func (u *userToken) isNotExpired(token string, expireAtSec int64) (*my_jwt.CustomClaims, int) {
	if customClaims, err := u.userJwt.ParseToken(token); err == nil {
		if time.Now().Unix()-(customClaims.ExpiresAt+expireAtSec) < 0 {
			// token 有效
			return customClaims, consts.JwtTokenOK
		} else {
			// 过期的 token
			return customClaims, consts.JwtTokenExpired
		}
	} else {
		// 无效的 token
		return nil, consts.JwtTokenInvalid
	}
}

// IsEffective 判断token是否有效（未过期+数据库用户信息正常）
func (u *userToken) IsEffective(token string) bool {
	customClaims, code := u.isNotExpired(token, 0)
	if consts.JwtTokenOK == code {
		// 1.首先在 redis 中检测是否存在某个用户对应的有效 token，如果存在就直接返回，不再继续查询 MySQL，否则最后查询 MySQL，确保万无一失
		tokenRedisFactory := token_cache_redis.CreateUsersTokenCacheFactory(customClaims.UID)
		if tokenRedisFactory != nil {
			defer tokenRedisFactory.ReleaseRedisConn()
			if tokenRedisFactory.TokenCacheIsExists(token) {
				return true
			}
		}
		// 2.token符合token本身的规则以后，继续在数据库校验是不是符合本系统其他设置，例如：一个用户默认只允许10个账号同时在线（10个token同时有效）
		if user.CreateUserFactory("").OauthCheckTokenIsOk(customClaims.UID, token) {
			return true
		}
	}
	return false
}

// RecordLoginToken 用户login成功，记录用户token
func (u *userToken) RecordLoginToken(userToken, clientIp string) bool {
	if customClaims, err := u.userJwt.ParseToken(userToken); err == nil {
		uid := customClaims.UID
		expiresAt := customClaims.ExpiresAt
		return user.CreateUserFactory("").OauthLoginToken(uid, userToken, expiresAt, clientIp)
	} else {
		return false
	}
}
