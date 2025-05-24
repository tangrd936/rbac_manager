package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"rbac_manager/common"
	"rbac_manager/global"
	"rbac_manager/middleware"
	"rbac_manager/utils/md5"
	"strings"
)

type Image struct {
}

var fileMap = map[string]struct{}{
	".jpg":  struct{}{},
	".png":  struct{}{},
	".jpeg": struct{}{},
	".webp": struct{}{},
}

func (i *Image) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		common.FailWithMsg(c, "请上传文件", err)
		return
	}
	// 判断文件大小是否大于2M
	if file.Size > 2*1024*1024 {
		common.FailWithMsg(c, "文件大小大于2M", nil)
		return
	}
	// 图片后缀判断
	suffix := path.Ext(file.Filename)
	_, ok := fileMap[strings.ToLower(suffix)]
	if !ok {
		common.FailWithMsg(c, "文件不合法", nil)
		return
	}

	auth := middleware.GetAuth(c)
	dst := fmt.Sprintf("static/image/%s/%s", auth.UserName, file.Filename)

	// 文件名重复判断,SaveUploadedFile方法会将重名文件覆盖
	_, err = os.Stat(dst)
	if err == nil {
		// 上传文件和之前文件重名了,使用文件hash判断文件是否相同
		newFile, _ := file.Open()
		newHash := md5.FileToMD5(newFile)

		oldFile, _ := os.Open(dst)
		oldHash := md5.FileToMD5(oldFile)
		if newHash != oldHash {
			// 文件内容不相同
			global.Log.Info(fmt.Sprintf("文件不相同, oldHash： %s; newHash: %s \n", oldHash, newHash))
			newName := strings.TrimRight(file.Filename, suffix) + "_" + newHash + suffix
			dst = fmt.Sprintf("static/image/%s/%s", auth.UserName, newName)
		}
	}
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		common.FailWithMsg(c, "文件保存失败", err)
		return
	}
	common.Ok(c, "/"+dst, "文件保存成功")
}
