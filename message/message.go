package message

import "encoding/xml"

//MsgType 基本消息类型
type MsgType string

//EventType 事件类型
type EventType string

//所有基本消息类型
const (
	MsgTypeText       MsgType = "text"                      // 文本消息
	MsgTypeImage              = "image"                     // 图片消息
	MsgTypeVoice              = "voice"                     // 音频消息
	MsgTypeVideo              = "video"                     // 视频消息
	MsgTypeShortVideo         = "shortvideo"                // 短视频消息[限接收=用户->微信服务]
	MsgTypeLocation           = "location"                  // 坐标消息/位置信息[限接收]
	MsgTypeLink               = "link"                      // 链接消息[限接收]
	MsgTypeMusic              = "music"                     // 音乐消息[限回复=微信服务->用户]
	MsgTypeNews               = "news"                      // 图文消息[限回复]
	MsgTypeTransfer           = "transfer_customer_service" // 消息转发到客服
)

// 所有事件类型
const (
	EventSubscribe   EventType = "subscribe"   //订阅
	EventUnsubscribe           = "unsubscribe" //取消订阅
	EventScan                  = "SCAN"        // 用户已经关注微信号，则微信会将带场景值的扫描事件推送给开发者
	EventLocation              = "LOCATION"    //上报地理位置事件
	EventClick                 = "CLICK"       // 点击菜单拉取消息时的事件推送
	EventView                  = "VIEW"        // 点击菜单跳转链接时的事件推送
)

// CommonToken 消息的通用结构部分
type CommonToken struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`   // 开发者微信号
	FromUserName string   `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64    `xml:"CreateTime"`   //	消息创建时间 （整型）
	MsgType      MsgType  `xml:"MagType"`      // 消息类型，文本为text
}

// 设置开发者微信号
func (msg *CommonToken) SetToUserName(toUserName string) {
	msg.ToUserName = toUserName
}

// 设置发送方帐号
func (msg *CommonToken) SetFromUserName(fromUserName string) {
	msg.FromUserName = fromUserName
}

// 设置消息创建时间
func (msg *CommonToken) SetCreateTime(createTime int64) {
	msg.CreateTime = createTime
}

// 设置消息类型
func (msg *CommonToken) SetMsgType(msgType MsgType) {
	msg.MsgType = msgType
}

// 存放所有微信发送过来的消息和事件
type MixMessage struct {
	CommonToken // 消息的通用结构

	//基本消息
	MsgID        int64   `xml:"MsgId"`        // 消息id，64位整型
	Content      string  `xml:"Content"`      // 文本消息内容
	PicURL       string  `xml:"PicUrl"`       // 图片链接（由系统生成）
	MediaID      string  `xml:"MediaId"`      // 图片/语音消息媒体id，可以调用获取临时素材接口拉取数据。
	Format       string  `xml:"Format"`       // 语音格式，如amr，speex等
	ThumbMediaID string  `xml:"ThumbMediaId"` // 视频消息缩略图的媒体id
	LocationX    float64 `xml:"Location_X"`   // 地理位置维度
	LocationY    float64 `xml:"Location_Y"`   // 地理位置经度
	Scale        float64 `xml:"Scale"`        // 地图缩放大小
	Label        string  `xml:"Label"`        // 地理位置信息
	Title        string  `xml:"Title"`        // 链接消息-消息标题
	Description  string  `xml:"Description"`  // 链接消息-消息描述
	URL          string  `xml:"Url"`          // 链接消息-消息链接

	// 事件相关
	Event     string `xml:"Event"`     //	事件类型
	EventKey  string `xml:"EventKey"`  // 扫描带参数二维码事件-事件KEY值
	Ticket    string `xml:"Ticker"`    // 二维码的ticket，可用来换取二维码图片
	Latitude  string `xml:"Latitude"`  //	地理位置纬度
	Longitude string `xml:"Longitude"` //	地理位置经度
	Precision string `xml:"Precision"` //	地理位置精度
}

// EncryptedXMLMsg 安全模式下的消息体
type EncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"` // 表示不进行序列化
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	EncryptedMsg string   `xml:"Encrypt" json:"Encrypt"`
}
