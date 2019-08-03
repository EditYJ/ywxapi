package ywxapi

import (
	"net/http"
	"ywxapi/context"
	"ywxapi/log"
	"ywxapi/server"
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
	channelLen := int64(10000)
	adapterName := "console"
	config := ""
	logLevel := log.LevelDebug
	log.InitLogger(channelLen, adapterName, config, logLevel)

	context := new(context.Context)
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

func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}
