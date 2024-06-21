package m_tikker

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/handy-golang/go-tools/m_encrypt"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_str"
)

func (obj *TikkerObj) RunToPm2() error {
	fileName := m_str.Join(
		"t_", m_encrypt.RandStr(6), ".sh",
	)

	Body := new(bytes.Buffer)
	Tmpl := template.Must(template.New("").Parse(TikkerSh))
	Tmpl.Execute(Body, TikkerShParam{
		Path:      obj.Path,
		FileName:  fileName,
		ShellCont: obj.Shell,
		LogPath:   obj.LogPath,
	})

	Cont := Body.String()
	filePath := m_str.Join(
		obj.Path,
		m_str.ToStr(os.PathSeparator),
		fileName,
	)

	m_file.Write(filePath, Cont)

	res, err := exec.Command("pm2", "start", filePath, "--name", fileName, "--no-autorestart").Output()
	if err != nil {
		errStr := fmt.Errorf("执行失败:%+v", m_str.ToStr(err))
		obj.Log.Println(errStr)
		return errStr
	} else {
		obj.Log.Println("执行成功", m_str.ToStr(res))
		return nil
	}
}

func (obj *TikkerObj) RunToShell() error {
	fileName := m_str.Join(
		"t_", m_encrypt.RandStr(3), ".sh",
	)

	Body := new(bytes.Buffer)
	Tmpl := template.Must(template.New("").Parse(TikkerSh))
	Tmpl.Execute(Body, TikkerShParam{
		Path:      obj.Path,
		FileName:  fileName,
		ShellCont: obj.Shell,
		LogPath:   obj.LogPath,
	})

	Cont := Body.String()
	shellPath := m_str.Join(
		obj.Path,
		m_str.ToStr(os.PathSeparator),
		fileName,
	)

	m_file.Write(shellPath, Cont)

	res, err := exec.Command("/bin/bash", shellPath).Output()
	if err != nil {
		errStr := fmt.Errorf("执行失败:%+v", m_str.ToStr(err))
		obj.Log.Println(errStr)
		return errStr
	} else {
		obj.Log.Println("执行成功", m_str.ToStr(res))
		return nil
	}
}
