package m_file

import "strings"

// 根据文件内容获取文件后缀名
func ContentToExtName(lType string) string {
	ext := ""
	switch lType {

	case "image/bmp":
		ext = "bmp"

	case "image/gif":
		ext = "gif"

	case "image/jpeg":
		ext = "jpeg"

	case "image/webp":
		ext = "webp"

	case "image/png":
		ext = "png"

	case "text/html":
		ext = "html"

	case "text/plain":
		ext = "txt"

	case "application/vnd.visio":
		ext = "vsd"

	case "application/vnd.ms-powerpoint":
		ext = "pptx"

	case "application/msword":
		ext = "docx"

	case "application/msexcel":
		ext = "xlsx"

	case "application/csv":
		ext = "csv"

	case "text/xml":
		ext = "xml"

	case "video/mp4":
		ext = "mp4"

	case "video/x-msvideo":
		ext = "avi"

	case "video/quicktime":
		ext = "mov"

	case "video/mpeg":
		ext = "mpeg"

	case "video/x-ms-wmv":
		ext = "wm"

	case "video/x-flv":
		ext = "flv"

	case "video/x-matroska":
		ext = "mkv"

	}

	if strings.Contains(lType, "text/html") {
		ext = "html"
	}

	return ext
}
