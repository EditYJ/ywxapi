package server

import (
	"fmt"
	"github.com/EditYJ/ywxapi/context"
	"github.com/EditYJ/ywxapi/util"
	"io/ioutil"
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

// 判断是否请求的确来自微信服务器
//
// 返回echostr给服务器表示接入成功
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

// 处理微信的请求判断是否为加密模式
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
	if srv.isSafeMode {
		// TODO 如果是安全模式，需要书写解密逻辑
	} else {
		rawXMLMsgBytes, err = ioutil.ReadAll(srv.Request.Body)
		if err != nil {
			return nil, fmt.Errorf("从bofy中解析XML失败，err=%v", err)
		}
	}
	srv.rawXMLMsg = string(rawXMLMsgBytes)
	fmt.Println(srv.rawXMLMsg)
	return nil, nil
}
