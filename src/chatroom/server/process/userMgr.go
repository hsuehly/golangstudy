package process

import "fmt"

// 因为在很多地方会使用，定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//  完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{onlineUsers: make(map[int]*UserProcess, 1024)}
}

// 完成对onlieUsers添加

func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserId] = up

}

//删除

func (um *UserMgr) DelOnlineUser(userId int) {
	delete(um.onlineUsers, userId)

}

// 返回当前所有在线的用户

func (um *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return um.onlineUsers

}

// 根据if返回对应的值

func (um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	// 如何冲map 取一个值
	up, ok := um.onlineUsers[userId]
	if !ok {
		// 说明查找的这个用户id不在线
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
