package m_time

import (
	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_str"
)

type GetTimeReturnType struct {
	TimeUnix int64  `bson:"TimeUnix"`
	TimeStr  string `bson:"TimeStr"`
}

func GetTime() (resData GetTimeReturnType) {
	resData.TimeUnix = GetUnixInt64()
	resData.TimeStr = UnixFormat(resData.TimeUnix)
	return
}

// ms 为 毫秒 时间戳
func TimeGet(ms any) (resData GetTimeReturnType) {
	myTime := MsToTime(ms, "0")
	resData.TimeUnix = ToUnixMsec(myTime)
	resData.TimeStr = UnixFormat(resData.TimeUnix)
	return
}

// 将毫秒时间戳转换为 x 时 x 分 x 秒

type HMSType struct {
	HH  string
	MM  string
	SS  string
	HMS string
}

func UnixTo_hh_mm_ss(ms any) HMSType {
	msToStr := m_str.ToStr(ms)

	H := 0
	M := 0
	S := "0"

	// 计算小时
	in_h := m_count.Div(msToStr, UnixTime.Hour)
	in_m := "0"
	in_s := "0"

	H = m_count.ToInt(in_h)

	h_dec := m_count.GetDecimal(in_h)
	if h_dec > 0 {
		// 计算分钟
		// 减去小时
		H_unix := UnixTimeInt64.Hour * int64(H)
		in_m_unix := m_count.Sub(msToStr, m_str.ToStr(H_unix))
		in_m = m_count.Div(in_m_unix, UnixTime.Minute)
	}
	M = m_count.ToInt(in_m)

	m_dec := m_count.GetDecimal(in_m)
	if m_dec > 0 {
		// 计算秒  减去 分钟 和 小时
		HM_unix := UnixTimeInt64.Hour*int64(H) + UnixTimeInt64.Minute*int64(M)
		in_s_unix := m_count.Sub(msToStr, m_str.ToStr(HM_unix))
		in_s = m_count.Div(in_s_unix, UnixTime.Seconds)
	}
	S = m_count.CentRound(in_s, 1)

	ReturnS := ""
	ReturnS = m_str.Join(ReturnS, H, "时")
	ReturnS = m_str.Join(ReturnS, M, "分")
	ReturnS = m_str.Join(ReturnS, S, "秒")

	return HMSType{
		HH:  m_str.ToStr(H),
		MM:  m_str.ToStr(M),
		SS:  S,
		HMS: ReturnS,
	}
}
