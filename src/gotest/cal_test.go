package gotest

import "testing"

func TestAdd(t *testing.T) {
	res := add(10, 10)
	if res != 20 {
		t.Fatalf("add执行错误，期望得到%v 实际得到%v\n", 20, res)
	}
	t.Logf("代码执行正确")
}
func TestSub(t *testing.T) {
	res := Sub(10, 5)
	if res != 5 {
		t.Fatalf("add执行错误，期望得到%v 实际得到%v\n", 5, res)
	}
	t.Logf("代码执行正确")
}
