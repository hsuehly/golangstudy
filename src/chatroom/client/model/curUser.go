package model

import (
	"main/chatroom/common/message"
	"net"
)

// 因为在客户端很多要使用
// 全局

type CurUser struct {
	Conn net.Conn
	message.User
}
