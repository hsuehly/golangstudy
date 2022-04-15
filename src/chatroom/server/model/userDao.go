package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"main/chatroom/common/message"
)

// 我们在服务器启动后就初始化一个userDao 实列
// 把它做成全局的变量，在需要的和redis 操作是，就直接使用变量

var (
	MyUserDao *UserDao
)

type UserDao struct {
	poll *redis.Pool
}

// 使用工厂模式创建一个userDao实列

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{poll: pool}
	return

}

// 根据用户id 返回一个user实力
func (ud *UserDao) getUserById(conn redis.Conn, id int) (user User, err error) {
	fmt.Println("id", id)
	res, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil {
		fmt.Println("err--", err)
		if err == redis.ErrNil {
			//表示在redis users中没有找到对应的id
			err = ERROR_USER_NOTEXISTS
			fmt.Println("err-", err)

		}
		return
	}
	fmt.Println("resres", res)
	// 这里我们把res反序列化成一个对象
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json格式化失败", err)
		return

	}
	fmt.Println("user-----", user)
	return
}

// 完成登录的校验
// 1.login 完成用户的验证
// 如果id 和pwd 都正确，则返回一个user实列
// 如果用户id和pwd 有错误，则返回一个错误的信息

func (ud *UserDao) Login(userId int, userPwd string) (user User, err error) {
	// 先从连接池里取出一个连接
	conn := ud.poll.Get()
	defer conn.Close()
	user, err = ud.getUserById(conn, userId)
	if err != nil {
		return
	}
	// 这时证明这个用户是获取到的
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return

	}
	return
}

//
func (ud *UserDao) Register(user *message.User) (err error) {
	// 先从userdao的连接池取出起一个连接
	conn := ud.poll.Get()
	defer conn.Close()
	fmt.Println("userID", user.UserId)
	_, err = ud.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	// 这时说明id 还没有

	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	// 写入数据库
	_, err = conn.Do("Hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存用户信息出错", err)
		return
	}
	return
}
