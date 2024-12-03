package upload

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/global/variable"
	"douyin-backend/app/http/controller/web"
	"douyin-backend/app/utils/files"
	"douyin-backend/app/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Cover struct {
}

func (c Cover) CheckParams(ctx *gin.Context) {
	tmpFile, err := ctx.FormFile(variable.ConfigYml.GetString("FileUploadSetting.UploadFileField")) //  file 是一个文件结构体（文件对象）
	var isPass bool
	//获取文件发生错误，可能上传了空文件等
	if err != nil {
		response.Fail(ctx, consts.FilesUploadFailCode, consts.FilesUploadFailMsg, err.Error())
		return
	}
	//超过系统设定的最大值：32M，tmpFile.Size 的单位是 bytes 和我们定义的文件单位M 比较，就需要将我们的单位*1024*1024(即2的20次方)，一步到位就是 << 20
	sizeLimit := variable.ConfigYml.GetInt64("FileUploadSetting.Size")
	if tmpFile.Size == 0 {
		response.Fail(ctx, consts.FilesUploadMoreThanMaxSizeCode, consts.FilesUploadIsEmpty, "")
		return
	}

	if tmpFile.Size > sizeLimit<<20 {
		response.Fail(ctx, consts.FilesUploadMoreThanMaxSizeCode, consts.FilesUploadMoreThanMaxSizeMsg+strconv.FormatInt(sizeLimit, 10)+"M", "")
		return
	}
	//不允许的文件mime类型
	if fp, err := tmpFile.Open(); err == nil {
		mimeType := files.GetFilesMimeByFp(fp)

		for _, value := range variable.ConfigYml.GetStringSlice("FileUploadSetting.AllowMimeType") {
			if strings.ReplaceAll(value, " ", "") == strings.ReplaceAll(mimeType, " ", "") {
				isPass = true
				break
			}
		}
		_ = fp.Close()
	} else {
		response.ErrorSystem(ctx, consts.ServerOccurredErrorMsg, "")
		return
	}
	//凡是存在相等的类型，通过验证，调用控制器
	if !isPass {
		response.Fail(ctx, consts.FilesUploadMimeTypeFailCode, consts.FilesUploadMimeTypeFailMsg, "")
	} else {
		(&web.UploadController{}).Cover(ctx)
	}
}
