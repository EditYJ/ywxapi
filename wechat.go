package ywxapi

import (
	"github.com/EditYJ/ywxapi/context"
	"github.com/EditYJ/ywxapi/server"
	"net/http"
)

type Wechat struct {
	Context *context.Context
}

// 使用者的配置
type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
}

//初始化
func NewWeChat(cfg *Config) *Wechat {
	// 创建上下文
	context := new(context.Context)
	// 将配置信息输入上下文
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}

// 安放使用者配置信息
func copyConfigToContext(cfg *Config, ctx *context.Context) {
	ctx.AppID = cfg.AppID
	ctx.AppSecret = cfg.AppSecret
	ctx.Token = cfg.Token
	ctx.EncodingAESkey = cfg.EncodingAESKey
}

// 接受原生的Request和ResponseWriter用于和微信服务器沟通联系，获取返回api服务
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}
