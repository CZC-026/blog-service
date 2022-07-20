package upload

import (
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

//获取文件名称
func GetFileName(name string) string {
	ext := GetFileExt(name) //获取文件后缀
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

//
func GetFileExt(name string) string {
	return path.Ext(name)
}

//获取文件保存地址，这里直接返回配置中的文件保存目录即可，
//也便于后续的调整。
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

//检查保存目录是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst) //获取文件的描述信息 FileInfo
	return os.IsNotExist(err)
}

//检查文件后缀是否包含在约定的后缀配置项中
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

//检查文件大小是否超出最大大小限制。
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

//检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

//该方法将会以传入的 os.FileMode 权限位去递归创建所需的所有目录结构，
//若涉及的目录均已存在，则不会进行任何操作，直接返回 nil。
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst) //创建目标地址的文件
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src) //通过 file.Open 方法打开源地址的文件，结合 io.Copy 方法实现两者之间的文件内容拷贝。
	return err
}
