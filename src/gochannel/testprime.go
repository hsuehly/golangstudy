package gochannel

import (
	"fmt"
	"time"
)

func TestPrime() {
	start := time.Now().Unix()
	//lable1:
	for num := 1; num <= 100000; num++ {
		//var jj = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//jj = false
				//fmt.Println("prime", num)
				break
				//break lable1
			}

		}
		//fmt.Println("num", num)
		//if flag {
		//	break
		//}

	}
	end := time.Now().Unix()
	fmt.Println("普通耗时：", end-start)
}
