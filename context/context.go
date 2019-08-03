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

func (ctx *Context) String(str string) error {
	ctx.Writer.WriteHeader(200)
	_, err := ctx.Writer.Write([]byte(str))
	return err
}

// GetQuery和Query()类似,他返回查询的值
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
