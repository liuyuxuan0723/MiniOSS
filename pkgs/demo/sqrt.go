package demo

import "fmt"

// 练习1：实现平方根
func Sqrt(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
