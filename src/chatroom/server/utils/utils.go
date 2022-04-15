package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"main/chatroom/common/message"
	"net"
)

// 将方法封装到结构体上

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (tf *Transfer) ReadPkg() (msg message.Message, err error) {
	//buf := make([]byte, 8096)

	fmt.Println("读取客户端发送的数据。。。")
	// conn.Read 在conn没有别关闭的情况下，才会阻塞
	// 如果客户端关闭了conn 则不会阻塞
	_, err = tf.Conn.Read(tf.Buf[:4])
	if err != nil {
		//err = errors.New("读取客户端信息失败")
		return
	}
	// 根据buf[:4] 转成一个uint32 类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(tf.Buf[0:4])
	// 根据pkglen 读取消息内容
	n, err := tf.Conn.Read(tf.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("conn.Read fail err")
		return
	}
	// 把paklen 反序列化成message.Message
	err = json.Unmarshal(tf.Buf[:pkgLen], &msg)
	if err != nil {
		//err = errors.New("反序列化失败")
		return
	}
	return
}

func (tf *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度给客户端
	var pkgLen uint32
	// 获取发送信息的长度
	pkgLen = uint32(len(data))
	fmt.Println("包的长度", pkgLen)
	// 创建一个byte切片
	//var buf [4]byte
	binary.BigEndian.PutUint32(tf.Buf[0:4], pkgLen)
	//fmt.Println("buf", buf)
	n, err := tf.Conn.Write(tf.Buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.write(bytes) fail", err)
		return
	}
	// 发送数据
	n, err = tf.Conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.write(bytes) fail", err)
		return
	}
	return

}
