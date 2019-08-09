package context

import "net/http"

type Context struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESkey string

	Writer  http.ResponseWriter
	Request *http.Request
}

func (ctx *Context) getAccessToken() {

}

// 响应数据
//
// 用来响应微信服务器的一些数据
func (ctx *Context) String(str string) error {
	ctx.Writer.WriteHeader(200)
	_, err := ctx.Writer.Write([]byte(str))
	return err
}

// 查询请求数据
//
// 主要用来查询微信服务器发送的对应[key]下的数据
func (ctx *Context) GetQuery(key string) (string, bool) {
	req := ctx.Request
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}

// 当退出的时候返回url查询的值
func (ctx *Context) Query(key string) string {
	value, _ := ctx.GetQuery(key)
	return value
}
