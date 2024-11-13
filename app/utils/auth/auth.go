package auth

import (
	"douyin-backend/app/global/variable"
	"douyin-backend/app/http/middleware/my_jwt"
	"github.com/gin-gonic/gin"
)

func GetUidFromToken(ctx *gin.Context) (uid int64) {
	userToken, exists := ctx.Get("userToken")
	if exists {
		uid = userToken.(my_jwt.CustomClaims).UID
	} else {
		uid = variable.ConfigYml.GetInt64("Token.JwtDefaultUid")
		variable.ZapLog.Error(ctx.ClientIP() + "userToken.UID not exists!")
	}
	return
}
