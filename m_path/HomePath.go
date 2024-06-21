package m_path

import (
	"os"
	"runtime"
)

// 获取系统的根路径
func GetHomePath() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// 获取当前的工作目录
func GetPwd() string {
	var appPath, _ = os.Getwd()
	return appPath
}
