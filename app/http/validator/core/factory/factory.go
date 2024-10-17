package factory

import (
	"douyin-backend/app/core/container"
	"douyin-backend/app/global/my_errors"
	"douyin-backend/app/global/variable"
	"douyin-backend/app/http/validator/core/interf"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Create 表单参数验证器工厂（请勿修改）
func Create(key string) func(context *gin.Context) {
	if value := container.CreateContainersFactory().Get(key); value != nil {
		if val, isOk := value.(interf.ValidatorInterface); isOk {
			return val.CheckParams
		} else {
			fmt.Println(val)
		}
	}
	variable.ZapLog.Error(my_errors.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}
