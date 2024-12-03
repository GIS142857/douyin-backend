package upload_file

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/global/my_errors"
	"douyin-backend/app/global/variable"
	"douyin-backend/app/model/video"
	"douyin-backend/app/utils/md5_encrypt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

func UploadVideo(context *gin.Context, savePath string) (r bool, finalSavePath interface{}) {
	// 获取上传文件
	file, _ := context.FormFile("file") // 假设上传字段名为 "file"

	// 生成全局唯一文件名
	if sequence := variable.SnowFlake.GetId(); sequence > 0 {
		// 保存视频文件
		saveFileName := fmt.Sprintf("%d%s", sequence, file.Filename)
		saveFileName = md5_encrypt.MD5(saveFileName) + path.Ext(saveFileName)
		videoFilePath := path.Join(savePath, saveFileName)

		// 确定封面文件名和路径
		saveCoverFileName := strings.TrimSuffix(saveFileName, ".mp4") + ".png"
		coverSavePath := variable.ConfigYml.GetString("FileUploadSetting.UploadRootPath") + variable.ConfigYml.GetString("FileUploadSetting.VideoCoverUploadFileSavePath")
		coverFilePath := path.Join(coverSavePath, saveCoverFileName)

		// 保存上传文件到视频路径
		if saveErr := context.SaveUploadedFile(file, videoFilePath); saveErr == nil {
			// 提取第一帧画面
			if extractCoverFrame(videoFilePath, coverFilePath) != nil {
				context.JSON(400, gin.H{"error": "Failed to extract video cover"})
				return false, nil
			}

			// 构建视频和封面的访问路径
			playAddr := variable.ConfigYml.GetString("FileUploadSetting.SourceUrlPrefix") + variable.ConfigYml.GetString("FileUploadSetting.VideoUploadFileSavePath") + saveFileName
			coverAddr := variable.ConfigYml.GetString("FileUploadSetting.SourceUrlPrefix") + variable.ConfigYml.GetString("FileUploadSetting.VideoCoverUploadFileSavePath") + saveCoverFileName

			// 获取描述、标签和隐私状态
			videoDesc := context.GetString(consts.ValidatorPrefix+"description") + context.GetString(consts.ValidatorPrefix+"tags")
			privateStatus := context.GetFloat64(consts.ValidatorPrefix + "private_status")

			// 插入数据库
			insertStatus := video.CreateVideoFactory("").InsertVideo(context, playAddr, videoDesc, coverAddr, int(privateStatus))
			if insertStatus {
				finalSavePath = gin.H{
					"playAddr":  playAddr,
					"coverAddr": coverAddr,
				}
				return true, finalSavePath
			}
		} else {
			variable.ZapLog.Error("文件保存出错：" + saveErr.Error())
		}
	} else {
		variable.ZapLog.Error("雪花算法生成 ID 失败")
	}

	return false, nil
}

func UploadAvatar(context *gin.Context, savePath string) (r bool, finnalSavePath interface{}) {
	//newSavePath, newReturnPath := generateYearMonthPath(savePath)
	//  1.获取上传的文件名(参数验证器已经验证完成了第一步错误，这里简化)
	file, _ := context.FormFile(variable.ConfigYml.GetString("FileUploadSetting.UploadFileField")) //  file 是一个文件结构体（文件对象）
	//  保存文件，原始文件名进行全局唯一编码加密、md5 加密，保证在后台存储不重复
	var saveErr error
	if sequence := variable.SnowFlake.GetId(); sequence > 0 {
		saveFileName := fmt.Sprintf("%d%s", sequence, file.Filename)
		saveFileName = md5_encrypt.MD5(saveFileName) + path.Ext(saveFileName)
		if saveErr = context.SaveUploadedFile(file, savePath+saveFileName); saveErr == nil {
			urlAddr := variable.ConfigYml.GetString("FileUploadSetting.SourceUrlPrefix") + variable.ConfigYml.GetString("FileUploadSetting.AvatarSmallUploadFileSavePath") + saveFileName
			insertStatus := video.CreateVideoFactory("").UpdateAvatar(context, urlAddr)
			//  上传成功,返回资源的相对路径，这里请根据实际返回绝对路径或者相对路径
			if insertStatus {
				finnalSavePath = gin.H{
					"urlAddr": urlAddr,
				}
			}
			return true, finnalSavePath
		}
	} else {
		saveErr = errors.New(my_errors.ErrorsSnowflakeGetIdFail)
		variable.ZapLog.Error("文件保存出错：" + saveErr.Error())
	}
	return false, nil

}

func UploadCover(context *gin.Context, savePath string) (r bool, finnalSavePath interface{}) {
	//  1.获取上传的文件名(参数验证器已经验证完成了第一步错误，这里简化)
	file, _ := context.FormFile(variable.ConfigYml.GetString("FileUploadSetting.UploadFileField")) //  file 是一个文件结构体（文件对象）
	//  保存文件，原始文件名进行全局唯一编码加密、md5 加密，保证在后台存储不重复
	var saveErr error
	if sequence := variable.SnowFlake.GetId(); sequence > 0 {
		saveFileName := fmt.Sprintf("%d%s", sequence, file.Filename)
		saveFileName = md5_encrypt.MD5(saveFileName) + path.Ext(saveFileName)
		if saveErr = context.SaveUploadedFile(file, savePath+saveFileName); saveErr == nil {
			urlAddr := variable.ConfigYml.GetString("FileUploadSetting.SourceUrlPrefix") + variable.ConfigYml.GetString("FileUploadSetting.CoverUploadFileSavePath") + saveFileName
			insertStatus := video.CreateVideoFactory("").UpdateCover(context, urlAddr)
			//  上传成功,返回资源的相对路径，这里请根据实际返回绝对路径或者相对路径
			if insertStatus {
				finnalSavePath = gin.H{
					"urlAddr": urlAddr,
				}
			}
			return true, finnalSavePath
		}
	} else {
		saveErr = errors.New(my_errors.ErrorsSnowflakeGetIdFail)
		variable.ZapLog.Error("文件保存出错：" + saveErr.Error())
	}
	return false, nil

}

// 文件上传可以设置按照 xxx年-xx月 格式存储
func generateYearMonthPath(savePathPre string) (string, string) {
	returnPath := variable.BasePath + variable.ConfigYml.GetString("FileUploadSetting.UploadFileReturnPath")
	curYearMonth := time.Now().Format("2006_01")
	newSavePathPre := savePathPre + curYearMonth
	newReturnPathPre := returnPath + curYearMonth
	// 相关路径不存在，创建目录
	if _, err := os.Stat(newSavePathPre); err != nil {
		if err = os.MkdirAll(newSavePathPre, os.ModePerm); err != nil {
			variable.ZapLog.Error("文件上传创建目录出错" + err.Error())
			return "", ""
		}
	}
	return newSavePathPre + "/", newReturnPathPre + "/"
}

// 使用 FFmpeg 提取视频第一帧
func extractCoverFrame(videoPath, coverPath string) error {
	// FFmpeg 命令：`ffmpeg -i input.mp4 -frames:v 1 output.png`
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-frames:v", "1", coverPath)

	// 执行命令并检查错误
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract frame using ffmpeg: %v", err)
	}
	return nil
}
