package m_tikker

import (
	"log"
	"os"

	"github.com/handy-golang/go-tools/m_path"
)

/*
使用 pm2  启动一个完全独立的进程，执行全独立的Bash脚本，哪怕主进程已经挂了，脚本 依然可以顺利被执行完毕。
一般用于服务的自动更新

	m_tikker.NewTikker(m_tikker.TikkerOpt{
		ShellContent: `
		mEcho "我爱你22"
		`,
	}).RunToPm2()

*/

type TikkerOpt struct {
	ShellContent string
	LogPath      string
}

type TikkerObj struct {
	Path    string
	Shell   string
	Log     *log.Logger
	LogPath string
}

func NewTikker(opt TikkerOpt) *TikkerObj {
	var obj TikkerObj

	// 生成脚本执行目录
	Path := m_path.Dir.Home
	isPath := m_path.Exists(Path)
	if !isPath {
		// 不存在则创建目录
		os.MkdirAll(Path, 0o777)
	}

	// 日志存放目录
	LogPath := m_path.Dir.App + "/mTikkerLogs"
	if len(opt.LogPath) > 0 {
		LogPath = opt.LogPath
	}

	// 创建日志目录
	isLogPath := m_path.Exists(LogPath)
	if !isLogPath {
		// 不存在则创建 logs 目录
		os.MkdirAll(LogPath, 0o777)
	}

	obj.Path = Path
	obj.Shell = opt.ShellContent

	// 创建日志文件
	file := LogPath + "/" + "mTikker.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o777)
	if nil != err {
		return nil
	}
	obj.Log = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	obj.Log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	obj.LogPath = file

	obj.Log.Println("欢迎使用 goTools !")

	return &obj
}
