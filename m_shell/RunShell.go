package m_shell

import (
	"bytes"
	_ "embed"
	"os"
	"os/exec"
	"text/template"

	"github.com/handy-golang/go-tools/m_file"
)

//go:embed FreedomShell.sh
var FreedomShell string

type FreedomShellParam struct {
	ShellContent string
}

func Run(ShellCont string) (resData []byte, resErr error) {
	Body := new(bytes.Buffer)
	Tmpl := template.Must(template.New("").Parse(FreedomShell))
	Tmpl.Execute(Body, FreedomShellParam{
		ShellContent: ShellCont,
	})
	Cont := Body.String()

	NowPah, _ := os.Getwd()

	ShellPath := NowPah + "/FreedomShell.sh"
	m_file.Write(ShellPath, Cont)

	res, err := exec.Command("/bin/bash", ShellPath).Output()
	if err != nil {
		resErr = err
	} else {
		resData = res
	}

	os.Remove(ShellPath)

	return
}
