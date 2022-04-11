package FamilyAccount

import "fmt"

type FamilyAccount struct {
	// 声明
	key  string
	loop bool
	// 定义账户的余额
	blance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	//收支的详情，用字符串来记录
	details string
	//
	flag bool
}

func (fa *FamilyAccount) showDetails() {
	fmt.Println("-------------当前收支明细记录-------------")
	if fa.flag {
		fmt.Println(fa.details)

	} else {
		fmt.Println("没有收支明细，来一笔吧。。。")

	}
}
func (fa *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&fa.money)
	fa.blance += fa.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&fa.note)
	fa.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", fa.blance, fa.money, fa.note)
	fa.flag = true
}
func (fa *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&fa.money)
	if fa.money > fa.blance {
		fmt.Println("余额不足")
		//fa.loop = false
		//return
	}
	fa.blance -= fa.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&fa.note)
	fa.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", fa.blance, fa.money, fa.note)
	fa.flag = true

}
func (fa *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("请输入正确的值 y/n")

	}
	if choice == "y" {
		fa.loop = false

	}
}
func NewFamAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    false,
		blance:  10000.0,
		money:   0.0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说   明",
		flag:    false,
	}
}
func (fa *FamilyAccount) MainMenu() {

	for {
		fmt.Println("\n-------------家庭收支记账软件-------------")
		fmt.Println(" 			1.收入明细")
		fmt.Println(" 			2.登记收入")
		fmt.Println(" 			3.登记支出")
		fmt.Println(" 			4.退出软件")
		fmt.Println("请选择(1-4): ")

		fmt.Scanln(&fa.key)
		switch fa.key {
		case "1":
			fa.showDetails()
		case "2":
			fa.income()
		case "3":
			fa.pay()
		case "4":
			//loop = false
			fa.exit()
		default:

			fmt.Println("请输入正确的选项。。。")
		}
		if !fa.loop {
			break
		}

	}
}
