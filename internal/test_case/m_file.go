package test_case

import "github.com/handy-golang/go-tools/m_file"

func Test_m_file() {

	m_file.WriteByte("/root/handy-golang/go-tools/data/111/222/333/444/555/test.txt", []byte("hello world"))

	// filetype := m_file.GetContentType("test.txt")

	// fmt.Println("filetype", filetype)

	// filecont := m_file.ReadFile("test.txt")

	// fmt.Println("filecont", m_str.ToStr(filecont))

	// fileName := m_file.GetSaveFileName(
	// 	m_file.GetNameOpt{
	// 		FileName: "test.txt",
	// 		SavePath: "./",
	// 		RandName: true,
	// 	},
	// )

	// fmt.Println("fileName", m_str.ToStr(fileName))

	// m_file.DownFile(
	// 	m_file.DownFileOpt{
	// 		Url: "http://file.mo7.cc/api/public/bz",
	// 	},
	// )

}
