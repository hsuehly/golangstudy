package view

import (
	"fmt"
	"main/customerMange/model"
	"main/customerMange/service"
)

type customerView struct {
	key  string
	loop bool
	// customerService
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (cv *customerView) list() {
	// 首先，获取到当前素有的客户信息（在切片中）
	customers := cv.customerService.List()
	fmt.Println("-----------客户列表-----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("\n-----------客户列表完成-----------\n")
}

// 得到新的用户，添加到切片中
func (cv *customerView) add() {
	fmt.Println("-----------添加客户-----------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	//service.CurrentEdit.Name=name
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	custoer := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(custoer) {
		fmt.Println("=============添加完成=============")
	} else {
		fmt.Println("=============添加失败=============")

	}
}
func (cv *customerView) delete() {
	fmt.Println("------------删除客户------------")
	fmt.Println("请选择待删除的客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N)")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if cv.customerService.Delete(id) {
			fmt.Println("=============删除完成=============")
		} else {
			fmt.Println("=============删除失败，输入id不存在=============")

		}

	}

}
func (cv *customerView) update() {

	fmt.Println("-----------更新客户-----------")
	fmt.Println("请输入更新客户的编号：")
	id := ""
	fmt.Scanln(&id)
	//strconv.ParseInt()

	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	//service.CurrentEdit.Name=name
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
}
func (cv *customerView) exit() {
	fmt.Println("确认是否退出(Y/N): ")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "y" || cv.key == "Y" || cv.key == "n" || cv.key == "N" {
			break
		}
		fmt.Println("你输入有误，确认是否退出(Y/N)")
	}
	if cv.key == "y" || cv.key == "Y" {
		cv.loop = false
	}

}
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("------------客户信息管理软件------------")
		fmt.Println("			 1. 添 加 客 户")
		fmt.Println("			 2. 修 改 客 户")
		fmt.Println("			 3. 删 除 客 户")
		fmt.Println("			 4. 客 户 列 表")
		fmt.Println("			 5. 退      出")
		fmt.Println("请选择(1-5): ")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			//fmt.Println("修 改 客 户")
			//service.CurrentEdit= cv.customerService.FindById(1)
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("请输入正确的值(1-5)")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("你退出了客户管理系统")
}
func CustomerView() {
	var cv = customerView{
		key:  "",
		loop: true,
	}
	// 这里完成对customerView 结构体的初始化工作
	cv.customerService = service.NewCustomerService()
	cv.mainMenu()
}
