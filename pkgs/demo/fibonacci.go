package demo

import "fmt"

// 闭包实现fibonacci
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		tmp := a
		a, b = b, a+b
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
