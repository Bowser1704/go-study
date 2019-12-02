package _interface

import "fmt"

type tester interface {
	test() string
}
type OldTester struct {
	info string
}
type NewTester struct {
	tester
	info string
}

func (older OldTester) test() string {
	return older.info
}
func (newer NewTester) test() string {
	return newer.info
}

func printTest(t tester) {
	fmt.Println(t.test())
}

func main() {
	older := OldTester{info: "old"}
	// 必须是实现了tester interface的类型才可以实例化
	newer := NewTester{tester: older, info: "new"}
	printTest(older) //old
	printTest(newer) //new
	//如果NewTester没有实现tester，下面也会输出old
}
