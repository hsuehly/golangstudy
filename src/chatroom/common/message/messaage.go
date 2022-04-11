package message

// 类型常量
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
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
	Code  int    `json:"code"`  // 返回的状态码 500 服务端错误， 200 ok
	Error string `json:"error"` // 返回错误信息
}
