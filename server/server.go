package server

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/EditYJ/ywxapi/context"
	"github.com/EditYJ/ywxapi/message"
	"github.com/EditYJ/ywxapi/util"
	"io/ioutil"
	"reflect"
	"strings"
)

type Server struct {
	*context.Context

	openID         string
	messageHandler func(message.MixMessage) *message.Reply

	requestRawXMLMsg []byte
	requestMsg       message.MixMessage

	responseRawXMLMsg []byte
	responseMsg       interface{}

	isSafeMode bool
	random     []byte
	nonce      string
	timestamp  int64
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
	// 校验
	if !srv.validate() {
		return fmt.Errorf("请求校验失败，该请求可能不是来自微信服务器！")
	}

	// 校验通过 发送echostr
	echostr, ok := srv.GetQuery("echostr")
	if ok {
		srv.String(echostr)
		return nil
	}

	// 生成回复消息
	response, err := srv.handleRequest()
	if err != nil {
		return err
	}

	srv.buildResponse(response)
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

// 处理微信的请求，根据微信请求的信息输出回复消息
// 此处利用用户定义的处理函数进行消息语义解析，生成回复消息
func (srv *Server) handleRequest() (reply *message.Reply, err error) {
	// 检查并设置是否为加密模式
	srv.isSafeMode = false
	encrypType := srv.Query("encrypt_type")
	if encrypType == "aes" {
		srv.isSafeMode = true
	}

	// 设置openID
	srv.openID = srv.Query("openid")

	// 接受微信请求消息并转换
	var msg interface{}
	msg, err = srv.getMessage()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	// 转换
	mixMessage, success := msg.(message.MixMessage)
	if !success {
		err = errors.New("消息类型转换失败")
	}
	// 保存微信请求消息
	srv.requestMsg = mixMessage
	// 利用用户自定函数处理微信请求消息/生成回复消息
	reply = srv.messageHandler(mixMessage)
	return
}

// 解析微信服务器返回的消息组装成对象
// TODO 增加对加密模式的支持(安全模式)
func (srv *Server) getMessage() (interface{}, error) {
	var rawXMLMsgBytes []byte
	var err error

	// 判断是否为加密模式
	if srv.isSafeMode {
		// TODO 如果是安全模式，需要书写解密逻辑
	} else {
		// 取出微信服务器请求的body内容(XML) 放入rawXMLMsgBytes
		rawXMLMsgBytes, err = ioutil.ReadAll(srv.Request.Body)
		if err != nil {
			return nil, fmt.Errorf("从body中解析XML失败，err=%v", err)
		}
	}

	// 原始消息存入[srv.requestRawXMLMsg]
	srv.requestRawXMLMsg = rawXMLMsgBytes
	// 返回解析后的对象
	return srv.parseRequestMessage(rawXMLMsgBytes)
}

// 将XML解析成对象
func (srv *Server) parseRequestMessage(rawXMLMsgBytes []byte) (msg message.MixMessage, err error) {
	msg = message.MixMessage{}
	err = xml.Unmarshal(rawXMLMsgBytes, &msg)
	return
}

// 设置用户自定义的回调方法
func (srv *Server) SetMessageHandler(handler func(message.MixMessage) *message.Reply)  {
	srv.messageHandler = handler
}

func (srv *Server) buildResponse(reply *message.Reply) (err error) {
	defer func() {
		if e:=recover(); e!=nil{
			err = fmt.Errorf("panic error: %v", err)
		}
	}()

	if reply == nil{
		// TODO 生成的回复消息如果为空，处理一下
		return nil
	}

	// 根据不一样的消息类型做出不一样的处理
	msgType := reply.MsgType
	switch msgType {
	case message.MsgTypeText:
	case message.MsgTypeImage:
	case message.MsgTypeVoice:
	case message.MsgTypeVideo:
	case message.MsgTypeMusic:
	case message.MsgTypeNews:
	case message.MsgTypeTransfer:
	default:
		err = message.ErrUnsupportReply
		return
	}

	// 因为不知道[reply.MsgData]的类型，所以需要判断[reply.MsgData]类型
	//
	// reflect.ValueOf示例:
	// var x int = 1
	// fmt.Println("value: ", reflect.ValueOf(x))
	// value:  <int Value>
	msgData := reply.MsgData
	value := reflect.ValueOf(msgData)
	// 取出类型，规定类型必须为“ptr”(指针类型)
	kind := value.Kind().String()
	if 0 != strings.Compare("ptr", kind){
		return message.ErrUnsupportReply
	}

	// 因为不知道reply.MsgData的具体类型(其实这个类型我们可以断定它是[CommonToken]的组合对象)所以需要
	// 反射调用[reply.MsgData]下的[SetToUserName],[SetFromUserName],[SetMsgType],[SetCreateTime]方法
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(srv.requestMsg.FromUserName)
	value.MethodByName("SetToUserName").Call(params)

	params[0] = reflect.ValueOf(srv.requestMsg.ToUserName)
	value.MethodByName("SetFromUserName").Call(params)

	params[0] = reflect.ValueOf(msgType)
	value.MethodByName("SetMsgType").Call(params)

	params[0] = reflect.ValueOf(util.GetCurrTs())
	value.MethodByName("SetCreateTime").Call(params)

	srv.responseMsg = msgData
	srv.responseRawXMLMsg, err = xml.Marshal(msgData)
	return
}
