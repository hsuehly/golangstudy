package zhizen

import "fmt"

func ZhiFc() {
	var h int = 100
	var f *int = &h

	fmt.Println(&h, &f)
}
