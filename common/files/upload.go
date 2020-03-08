package files

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"sisyphus/common/setting"
	"sisyphus/common/utils"
	"strings"
)

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return setting.GetAppConf().PrefixUrl + "/" + GetImagePath() + name
}

func GetImageFullPath() string {
	fmt.Println(path.Join(setting.GetAppConf().StaticRootPath, GetImagePath()))
	return path.Join(setting.GetAppConf().StaticRootPath, GetImagePath())
}

func GetImagePath() string {
	return setting.GetAppConf().ImageSavePath
}

func GenImageName(name string) string {
	ext := path.Ext(name)
	filename := strings.TrimSuffix(name, ext)
	filename = utils.EncodeMD5(filename)
	return filename + ext
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := GetExt(fileName)
	for _, allowExt := range setting.GetAppConf().ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImage check if the file exists
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}

// CheckImageSize check image size
func CheckImageSize(f multipart.File) bool {
	size, err := GetSize(f)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
		return false
	}

	return size <= setting.GetAppConf().ImageMaxSize*1024*1024
}
