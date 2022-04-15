package message

// User 为了徐磊话和反序列化成功，我们必须保证
// 用户的信息json字符串key 和结构体的字段对应的tag名字一致！！！
type User struct {
	UserId     int    `json:"user_id"`
	UserPwd    string `json:"user_pwd"`
	UserName   string `json:"user_name"`
	UserStatus int    `json:"user_status"` // 用户状态
	Sex        string `json:"sex"`
}
