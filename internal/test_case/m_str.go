package test_case

import (
	"fmt"
	"go-tools/m_str"
)

func Test_m_str() {

	toStr()

	joinStr()

	temp()
}

func toStr() {
	// ===== 任意类型转字符串
	// var a = []rune("mo7欢迎你")
	// var a = []byte("mo7欢迎你")
	// var a = 10.97
	// var a = os.PathSeparator
	// var a = map[string]any{
	// 	"name": "张三",
	// 	"age":  24,
	// 	"sex":  '男',
	// }
	var a = []int32{1, 2, 3, 4, 5}
	str := m_str.ToStr(a)
	fmt.Println("str: ", str)
}

func joinStr() {
	// 混合拼接字符串
	var a = []int32{1, 2, 3, 4, 5}
	joinStr := m_str.Join("mo7", "欢迎你", a, "张三")
	fmt.Println("joinStr: ", joinStr)
}

func temp() {
	// 字符串模板
	var config = `
app.name = ${appName}
app.ip = ${appIP}
app.port = ${appPort}
`

	var dev = map[string]string{
		"appName": "my_ap123p",
		"appIP":   "0.0.0.0",
		"appPort": "8080",
	}
	s := m_str.Temp(config, dev)

	fmt.Println("temp", s)
}
