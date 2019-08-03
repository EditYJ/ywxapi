package util

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"sort"
)

// 将接受到的字符串数组进行字典排序，根据官网文档可接收到[token,timestamp,nonce],
// 将这三个key中的value组成切片放入此函数进行字典排序，用sha1算法加密得到[hashcode]
//
// hash文档
// 参考示例：https://www.cnblogs.com/wanghui-garcia/p/10452428.html
//type Hash interface {
//	// 通过嵌入的匿名io.Writer接口的Write方法向hash中添加更多数据，永远不返回错误
//	io.Writer
//	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
//	Sum(b []byte) []byte
//	// 重设hash为无数据输入的状态
//	Reset()
//	// 返回Sum会返回的切片的长度
//	Size() int
//	// 返回hash底层的块大小；Write方法可以接受任何大小的数据，
//	// 但提供的数据是块大小的倍数时效率更高
//	BlockSize() int
//}
func Signature(params ...string) string {
	// 字符串字典排序，先比较高位，相同的再比较低位
	sort.Strings(params)

	// 得到sha1实例
	h := sha1.New()

	// 对切片进行循环遍历，持续将切片中的每一个字符串添加到h中
	// 主要是hash对象都实现了[io.writer]接口
	for _,s:=range params{
		_, err := io.WriteString(h, s)
		if err != nil{
			log.Fatal(err)
		}
	}
	// 返回解析结果
	return fmt.Sprintf("%x", h.Sum(nil))
}