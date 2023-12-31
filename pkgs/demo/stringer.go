package demo

import "fmt"

// 练习5:实现Stringer接口 自定义对象输出格式
type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
