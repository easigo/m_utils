package m_encrypt

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_str"
)

// 获取 UUID
func GetUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}

// 获取timeID
func TimeID() string {
	now := time.Now().Format("20060102150405")
	uuid := GetUUID()
	uuidArr := strings.Split(uuid, "")
	start := m_count.GetRound(0, int64(len(uuidArr)-3))

	sArr := uuidArr[start : start+3]

	s := strings.Join(sArr, "")

	returnStr := m_str.Join(
		"m", s, now,
	)

	return returnStr
}
