package message

import "errors"

// 无效的回复
var ErrInvlidReply = errors.New("无效的回复消息")

// 不支持的回复类型
var ErrUnsupportReply = errors.New("不支持这个回复消息类型")

// 消息回复
type Reply struct {
	MsgType MsgType
	MsgData interface{}
}