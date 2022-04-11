package service

import "main/customerMange/model"

type CustomerService struct {
	customers []model.Customer
	// 晟声明一个字段，表示当前切片含有多少个客户
	customerNum int
}

//var CurrentEdit *model.Customer = nil

// 编写一个方法，可以返回*customerService

func NewCustomerService() *CustomerService {
	// 初始化切片
	customerService := &CustomerService{
		customerNum: 1,
	}
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@hsuhely.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户的切片

func (cs *CustomerService) List() []model.Customer {
	return cs.customers

}

// 添加客户到customers 切片

func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

// 根据id 删除客户

func (cs *CustomerService) Delete(id int) bool {
	index := cs.FinedById(id)
	if index == -1 {
		return false
	}
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

// 根据id 更新客户信息

func (cs *CustomerService) Update(id int) bool {
	index := cs.FinedById(id)
	if index == -1 {
		return false
	}
	//cs.customers[index] =
	return true
}

// 根据id 查找客户在切片中的对应下标

func (cs *CustomerService) FinedById(id int) int {
	index := -1
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			index = i
		}

	}
	return index

}
func (cs *CustomerService) FindIndexById(id int) int {
	index := -1
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			index = i
			break
		}

	}
	return index

}
func (cs *CustomerService) FindById(id int) *model.Customer {
	index := -1
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			index = i
			break
		}

	}
	return &cs.customers[index]

}
