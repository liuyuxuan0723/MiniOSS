package demo

import "fmt"

// 练习6: 自定义异常处理 抛出ErrNegativeSqrt异常
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt1(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt1(2))
	fmt.Println(Sqrt1(-2))
}
