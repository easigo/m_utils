package m_file

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/handy-golang/go-tools/m_path"
)

// 写入文件内容 fileName 为文件的路径
func Write(filePath string, content string) (resErr error) {
	resErr = nil
	dir, _ := filepath.Split(filePath) //  获取目录
	isExist := m_path.IsExist(dir)
	if !isExist {
		os.MkdirAll(dir, os.ModePerm)
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o777)
	if err != nil {
		resErr = err
	} else {
		n, _ := f.Seek(0, io.SeekCurrent)
		f.WriteAt([]byte(content), n)
		defer f.Close()
	}
	return
}

func WriteByte(filePath string, content []byte) (resErr error) {
	resErr = nil
	dir, _ := filepath.Split(filePath) //  获取目录
	isExist := m_path.IsExist(dir)
	if !isExist {
		os.MkdirAll(dir, os.ModePerm)
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o777)
	if err != nil {
		resErr = err
	} else {
		n, _ := f.Seek(0, io.SeekCurrent)
		f.WriteAt(content, n)
		defer f.Close()
	}
	return
}

// 获取一个文件的类型
func GetContentType(fileName string) string {
	if !m_path.IsFile(fileName) {
		// fmt.Errorf("fileName必须为一个文件")
		return ""
	}

	file, err := os.Open(fileName)
	if err != nil {
		return ""
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	bs := make([]byte, 512)
	_, err = file.Read(bs)
	if err != nil {
		return "application/octet-stream"
	}

	// // Use the net/http package's handy DectectContentType function. Always returns a valid
	// // content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(bs)

	return contentType
}

func ReadFile(fileName string) []byte {
	if !m_path.IsFile(fileName) {
		errStr := fmt.Errorf("fileName必须为一个文件")
		log.Println(errStr)
		return []byte("")
	}

	f, err := os.ReadFile(fileName)
	if err != nil {
		return []byte("")
	}
	return f
}
