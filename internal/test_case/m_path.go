package test_case

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_path"
)

func Test_m_path() {

	// str := m_path.Exists("/Users/meichangliang/meichangliang/handy-golang/go-tools")
	// str := m_path.Exists("/Users/meichangliang/meichangliang/hanxdy-golang/go-tools")
	// str := m_path.Exists("m_path/Exists.go")
	// str := m_path.Exists("/m_path/Exists.go")
	// str := m_path.Exists("m_path/Exists.gxo")

	// str := m_path.GetHomePath()

	str := m_path.GetPwd()

	fmt.Println("str: ", str)

}
