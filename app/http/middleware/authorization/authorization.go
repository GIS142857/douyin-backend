package authorization

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/global/variable"
	userstoken "douyin-backend/app/service/users/token"
	"douyin-backend/app/utils/response"
	"github.com/gin-gonic/gin"
)

type HeaderParams struct {
	Token string `header:"Token" binding:"required,min=20"`
}

// CheckTokenAuth 检查token完整性、有效性中间件
func CheckTokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		headerParams := HeaderParams{}

		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := ctx.ShouldBindHeader(&headerParams); err != nil {
			response.TokenErrorParam(ctx, consts.JwtTokenMustValid+err.Error())
			return
		}
		token := headerParams.Token
		if len(token) >= 20 {
			tokenIsEffective := userstoken.CreateUserFactory().IsEffective(token)
			if tokenIsEffective {
				if customeToken, err := userstoken.CreateUserFactory().ParseToken(token); err == nil {
					key := variable.ConfigYml.GetString("Token.BindContextKeyName")
					// token 验证通过并绑定在请求的上下文中
					ctx.Set(key, customeToken)
				} else {
					response.TokenParseFail(ctx, err)
				}
				ctx.Next()
			} else {
				response.ErrorTokenAuthFail(ctx)
			}
		} else {
			response.ErrorTokenBaseInfo(ctx)
		}
	}
}
