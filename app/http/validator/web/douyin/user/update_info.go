package user

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/http/controller/web"
	"douyin-backend/app/http/validator/core/data_transfer"
	"douyin-backend/app/utils/response"
	"github.com/gin-gonic/gin"
)

type UpdateInfo struct {
	OperationType
	Data
}

func (u UpdateInfo) CheckParams(ctx *gin.Context) {
	if err := ctx.ShouldBind(&u); err != nil {
		response.ValidatorError(ctx, err)
		return
	}
	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := data_transfer.DataAddContext(u, consts.ValidatorPrefix, ctx)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(ctx, "update-info 表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器，写具体业务逻辑
		(&web.UserController{}).UpdateInfo(extraAddBindDataContext)
	}
}
