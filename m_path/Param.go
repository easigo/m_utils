package m_path

import "os"

var appPath, _ = os.Getwd()

type DirType struct {
	Home string // Home 根目录
	App  string // APP 根目录
}

var Dir = DirType{
	Home: GetHomePath(),
	App:  appPath,
}
