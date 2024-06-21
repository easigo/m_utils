package m_tikker

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"

	"github.com/handy-golang/go-tools/m_encrypt"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_path"
	"github.com/handy-golang/go-tools/m_str"
)

// pm2 安装
func (obj *TikkerObj) InstPm2() *TikkerObj {
	filePath := m_path.Dir.App
	fileName := m_str.Join(
		"i_", m_encrypt.RandStr(5), ".sh",
	)

	Body := new(bytes.Buffer)
	Tmpl := template.Must(template.New("").Parse(InstPm2))
	Tmpl.Execute(Body, InstPm2Param{
		Path:     filePath,
		FileName: fileName,
	})
	Cont := Body.String()

	shellPath := m_str.Join(
		filePath,
		m_str.ToStr(os.PathSeparator),
		fileName,
	)

	m_file.Write(shellPath, Cont)

	res, err := exec.Command("/bin/bash", shellPath).Output()
	if err != nil {
		obj.Log.Println("环境安装失败", m_str.ToStr(err))
	} else {
		obj.Log.Println("环境安装成功", m_str.ToStr(res))
	}

	return obj
}
