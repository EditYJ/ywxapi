package message

// EncryptedXMLMsg 安全模式下的消息体
type EncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"` // 表示不进行序列化
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	EncryptedMsg string   `xml:"Encrypt" json:"Encrypt"`
}
