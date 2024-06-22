package m_encrypt

import (
	"math/rand"
	"time"
)

// [a-z0-9] 中随机一个字符串
func RandStr(length int) string {
	baseStr := "0123456789abcdefghijklmnopqrstuvwxyz"

	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	bytes := make([]byte, length)
	l := len(baseStr)
	for i := 0; i < length; i++ {
		bytes[i] = baseStr[r.Intn(l)]
	}
	return string(bytes)
}
