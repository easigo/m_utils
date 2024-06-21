package m_file

import (
	"os"
	"path"
	"strings"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_encrypt"
	"github.com/handy-golang/go-tools/m_path"
	"github.com/handy-golang/go-tools/m_str"
)

type GetNameOpt struct {
	FileName string
	SavePath string
	RandName bool
}

type GetNameType struct {
	Count    string
	Name     string
	SrcName  string
	ExtName  string
	SavePath string
}

// 获取当前可用的文件名
func GetSaveFileName(opt GetNameOpt) string {
	extName := path.Ext(opt.FileName)                      // 后缀名
	name := strings.Replace(opt.FileName, extName, "", -1) // 把后缀名换成空字符串

	if opt.RandName {
		name = m_encrypt.TimeID()
	}

	var Obj GetNameType
	Obj.Count = "0"
	Obj.Name = name
	Obj.SrcName = name
	Obj.ExtName = extName
	Obj.SavePath = opt.SavePath

	isThere := Obj.FileThere()

	if isThere {
		Obj.GetNewName()
	}

	return m_str.Join(
		Obj.Name, Obj.ExtName,
	)
}

func (obj *GetNameType) FileThere() bool {
	filePath := m_str.Join(
		obj.SavePath,
		m_str.ToStr(os.PathSeparator),
		obj.Name,
		obj.ExtName,
	)

	isFilePath := m_path.Exists(filePath)

	return isFilePath
}

func (obj *GetNameType) GetNewName() *GetNameType {
	obj.Count = m_count.Add(obj.Count, "1")
	obj.Name = m_str.Join(
		obj.SrcName, "_", obj.Count,
	)

	isThere := obj.FileThere()
	if isThere {
		return obj.GetNewName()
	}

	return obj
}
