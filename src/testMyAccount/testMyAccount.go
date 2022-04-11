package testMyAccount

import "fmt"

func TestMyAccount() {
	key := ""
	//loop := true
	// 定义账户的余额
	blance := 10000.0
	// 每次收支的金额
	money := 0.0
	// 每次收支的说明
	note := ""
	//收支的详情，用字符串来记录
	details := "收支\t账户金额\t收支金额\t说   明"
	//
	flag := false
label:

	for {
		fmt.Println("\n-------------家庭收支记账软件-------------")
		fmt.Println(" 			1.收入明细")
		fmt.Println(" 			2.登记收入")
		fmt.Println(" 			3.登记支出")
		fmt.Println(" 			4.退出软件")
		fmt.Println("请选择(1-4): ")

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("-------------当前收支明细记录-------------")
			if flag {
				fmt.Println(details)

			} else {
				fmt.Println("没有收支明细，来一笔吧。。。")

			}
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			blance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", blance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > blance {
				fmt.Println("余额不足")
				break
			}
			blance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", blance, money, note)
			flag = true

		case "4":
			//loop = false
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
				break label

			}
		default:

			fmt.Println("请输入正确的选项。。。")
		}
		//if !loop {
		//	break
		//}

	}
	fmt.Println("你退出了家庭记账软件的使用。。。")

}
