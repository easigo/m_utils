package m_path

import (
	"fmt"
	"os"
)

// 判断目录或文件是否存在（终将被废弃）
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(nil) {
		return false
	}
	fmt.Println("warning:请改用 m_path.Exists 终将被废弃,请改用 m_path.IsExist() ")
	return false
}

// 判断目录或文件是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(nil) {
		return false
	}
	return false
}
