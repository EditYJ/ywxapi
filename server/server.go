package server

import (
	"fmt"
	"io/ioutil"
	"ywxapi/context"
	"ywxapi/util"
)

type Server struct {
	*context.Context
	isSafeMode bool
	rawXMLMsg  string
}

func NewServer(context *context.Context) *Server {
	srv := new(Server)
	srv.Context = context
	return srv
}

// 处理微信的请求消息
func (srv *Server) Serve() error {
	if !srv.validate() {
		return fmt.Errorf("请求校验失败，该请求可能不是来自微信服务器！")
	}
	echostr, ok := srv.GetQuery("echostr")
	if ok {
		return srv.String(echostr)
	}
	srv.handleRequest()
	return nil
}

// 校验请求是否合法
func (srv *Server) validate() bool {
	//取出Token，timestamp，nonce加密与signature进行比较
	timestamp := srv.Query("timestamp")
	nonce := srv.Query("nonce")
	signature := srv.Query("signature")
	return signature == util.Signature(srv.Token, timestamp, nonce)
}

// 处理微信的请求
func (srv *Server) handleRequest() {
	srv.isSafeMode = false
	encrypType := srv.Query("encrypt_type")
	if encrypType == "aes" {
		srv.isSafeMode = true
	}
	_, err := srv.getMessage()
	if err != nil {
		fmt.Printf("%v", err)
	}
}

// 获取微信服务器返回的消息
func (srv *Server) getMessage() (interface{}, error) {
	var rawXMLMsgBytes []byte
	var err error
	if srv.isSafeMode{
		// TODO 如果是安全模式，需要书写解密逻辑
	}else{
		rawXMLMsgBytes, err = ioutil.ReadAll(srv.Request.Body)
		if err != nil{
			return nil, fmt.Errorf("从bofy中解析XML失败，err=%v", err)
		}
	}
	srv.rawXMLMsg = string(rawXMLMsgBytes)
	fmt.Println(srv.rawXMLMsg)
	return nil, nil
}
