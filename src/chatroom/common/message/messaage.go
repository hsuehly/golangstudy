package message

// 类型常量
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// 用户在线状态常量
const (
	UserOnline = iota
	Useroffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"`
}
type LoginMes struct {
	UserId   int    `json:"user_id"`   // 用户id
	UserPwd  string `json:"user_pwd"`  // 用户密码
	UserName string `json:"user_name"` // 用户名

}
type LoginResMes struct {
	Code    int `json:"code"` // 返回的状态码 500 服务端错误， 200 ok
	UserIds []int
	Error   string `json:"error"` // 返回错误信息
}
type RegisterMes struct {
	User User `json:"user"` // 类型就是User结构体
}
type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回的状态码 400 表示该用户已经占用 200 表示成功
	Error string `json:"error"` // 返回错误信息
}

// 配合服务器端推送新定义

type NotifyUserStatusMes struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}

// 增加一个smsmes
// 发送消息

type SmsMes struct {
	Conntent string `json:"conntent"`
	User            // 匿名结构体 继承
}
