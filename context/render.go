package context

import (
	"encoding/xml"
	"net/http"
)

//xml
var xmlContentType = []string{"application/xml; charset=utf-8"}
//text/html的意思是将文件的content-type设置为text/html的形式，浏览器在获取到这种文件时会自动调用html的解析器对文件进行相应的处理。
//text/plain的意思是将文件设置为纯文本的形式，浏览器在获取到这种文件时并不会对其进行处理。
var plainContentType = []string{"text/plain; charset=utf-8"}

// 转换string呈现
func (ctx *Context) String(str string)  {
	writeContextType(ctx.Writer, plainContentType)
	ctx.Render([]byte(str))
}

// 转换xml呈现
func (ctx *Context) XML(obj interface{})  {
	writeContextType(ctx.Writer, xmlContentType)
	bytes, err := xml.Marshal(obj)
	if err!=nil{
		panic(err)
	}
	ctx.Render(bytes)
}

// 传入字节信息写入输入流发送至微信服务器
func (ctx *Context) Render(bytes []byte) {
	ctx.Writer.WriteHeader(200)
	_,err :=ctx.Writer.Write(bytes)
	if err != nil{
		panic(err)
	}
}

func writeContextType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val:= header["Content-Type"]; len(val)==0{
		header["Content-Type"] = value
	}
}