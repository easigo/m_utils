package m_fetch

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/handy-golang/go-tools/m_url"
	jsoniter "github.com/json-iterator/go"
)

/*

	data := map[string]any{
		"op": "subscribe",
		"args": []string{
			"123", "345", "mei",
		},
	}

	resData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin: "http://localhost:9000",
		Path:   "/api/ping",
		Data:   data,
		Header: map[string]string{
			"Content-Type":  "appli1arset=utf-8",
			"Content-Type1": "appl2et=utf-8",
			"Content-Type2": "applicati3et=utf-8",
			"Content-Type3": "application4t=utf-8",
		},
	}).Get()
	if err != nil {
		fmt.Println("err", err)
	}
	jsonStr := mJson.JsonFormat(resData)
	fmt.Println("resData", jsonStr)

*/

type HttpOpt struct {
	Origin    string
	Path      string
	Data      []byte
	DataMap   map[string]any // 如果存在 DataMap 则会忽略  Data 参数内容
	Header    map[string]string
	Event     func(string, any) // s1 = succeed , err
	ProxyURLs []string          //  ["socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338"]
}

type Http struct {
	Url    string
	Data   []byte
	Header map[string]string
	Event  func(string, any)
	C      *colly.Collector
}

func NewHttp(opt HttpOpt) (_this *Http) {
	_this = &Http{}
	// 检查参数
	errStr := []string{}
	switch {
	case len(opt.Origin) < 1:
		errStr = append(errStr, "Origin")
		fallthrough
	case len(opt.Path) < 1:
		errStr = append(errStr, "Path")
	}
	if len(errStr) > 0 {
		errStr := fmt.Errorf("缺少参数:%+v", errStr)
		panic(errStr)
	}

	_this.Url = m_str.Join(opt.Origin, opt.Path)

	_this.Data = opt.Data
	if len(opt.DataMap) > 0 {
		_this.Data = m_json.ToJson(opt.DataMap)
	}

	_this.Header = opt.Header
	_this.Event = opt.Event
	if _this.Event == nil {
		_this.Event = func(s1 string, s2 any) {}
	}

	// 创建 colly 对象
	_this.C = colly.NewCollector()
	if len(opt.ProxyURLs) > 0 {
		// 设置代理
		rp, err := proxy.RoundRobinProxySwitcher(
			opt.ProxyURLs...,
		)

		if err != nil {
			_this.Event("proxy err", err)
		}
		_this.C.SetProxyFunc(rp)
	}
	_this.C.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json; charset=utf-8")
		r.Headers.Set("User-Agent", "go-tools - github.com/handy-golang/go-tools")
		// 添加header头
		for key, val := range _this.Header {
			r.Headers.Set(key, val)
		}
	})

	return
}

// 处理 Get 参数
func (_this *Http) disposeGetParam() *Http {
	urlO := m_url.InitUrl(_this.Url)

	var dataUn map[string]any
	jsoniter.Unmarshal(_this.Data, &dataUn)

	for key, val := range dataUn {
		v := fmt.Sprintf("%+v", val)
		urlO.AddParam(key, v)
	}
	_this.Url = urlO.String()
	return _this
}

// GET
func (_this *Http) Get() (resData []byte, resErr error) {
	// 处理 Get 参数
	_this.disposeGetParam()

	_this.C.OnResponse(func(r *colly.Response) {
		resData = r.Body
		_this.Event("succeed", resData)
	})
	_this.C.OnError(func(r *colly.Response, errStr error) {
		resData = r.Body
		resErr = errStr
		_this.Event("err", errStr)
	})

	_this.C.Visit(_this.Url)

	return
}

// Post
func (_this *Http) Post() (resData []byte, resErr error) {

	_this.C.OnResponse(func(r *colly.Response) {
		resData = r.Body
		_this.Event("succeed", resData)
	})
	_this.C.OnError(func(r *colly.Response, err error) {
		resData = r.Body
		resErr = err
		_this.Event("err", err)
	})

	_this.C.PostRaw(_this.Url, _this.Data)

	return
}
