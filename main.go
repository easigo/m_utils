package main

import (
	_ "embed"

	"github.com/handy-golang/go-tools/internal/test_case"
)

// //go:embed package.json
// var AppPackage []byte

func main() {
	// m_json.PrintlnForByte(AppPackage)

	// 测试 m_str 库
	// test_case.Test_m_str()

	// 测试 m_count 库
	test_case.Test_m_count()

	// 测试 m_path 库
	// test_case.Test_m_path()

	// 测试 m_file 库
	// test_case.Test_m_file()
}
