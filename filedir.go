package utils

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func isExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func isFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

// 文件是否存在
func IsFileExist(f string) bool {
	if isExist(f) {
		return isFile(f)
	}
	return false
}

// 文件夹是否存在
func IsDirExist(f string) bool {
	if isExist(f) {
		return !isFile(f)
	}
	return false
}

// 获取文件名称
func GetFileName(dst string) string {
	filenameall := path.Base(dst)
	filesuffix := path.Ext(dst)
	return filenameall[0 : len(filenameall)-len(filesuffix)]
}

// 获取文件加密名称
func GetFileMD5Name(dst string, b ...byte) string {
	filenameall := path.Base(dst)
	filesuffix := path.Ext(dst)
	return StringMD5V(filenameall[0:len(filenameall)-len(filesuffix)], b...)
}

// 从图片文件获取base64编码
func FileToBase64Code(dst string) (string, error) {
	buf, err := ioutil.ReadFile(dst)
	return base64.StdEncoding.EncodeToString(buf), err
}

// base64图片编码保存到文件
func Base64CodeToFile(code, dst string) error {
	buf, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, buf, 0640)
}

// @param dst 文件路径+后缀名
// 保存文件并转换类型
func SaveMultiFileExt(dst string, file *multipart.FileHeader) error {
	buf, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer buf.Close()
	src, err := file.Open()
	if err != nil {
		return err
	}

	_, err = io.Copy(buf, src)
	return err
}
