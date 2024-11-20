package message

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/http/controller/web"
	"douyin-backend/app/http/validator/core/data_transfer"
	"douyin-backend/app/utils/response"
	"github.com/gin-gonic/gin"
)

type AllMsg struct {
}

func (a AllMsg) CheckParams(ctx *gin.Context) {
	// 基本的参数验证
	if err := ctx.ShouldBind(&a); err != nil {
		response.ValidatorError(ctx, err)
		return
	}
	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := data_transfer.DataAddContext(a, consts.ValidatorPrefix, ctx)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(ctx, "get_all_msg 表单验证器json化失败", "")
		return
	} else {
		// 验证完成，调用控制器，写具体业务逻辑
		(&web.MessageController{}).GetAllMsg(extraAddBindDataContext)
	}
}
