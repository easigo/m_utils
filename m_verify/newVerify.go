package m_verify

import (
	"go-tools/m_count"
	"go-tools/m_str"
)

func NewCode() string {
	code := m_count.GetRound(100000, 999999)
	return m_str.ToStr(code)
}
