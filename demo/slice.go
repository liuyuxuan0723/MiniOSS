package demo

// 练习2:返回一个切片输出指定函数图像
func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		slice[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			slice[i][j] = avg(i, j)
		}
	}
	return slice
}

func avg(x, y int) uint8 {
	return uint8((x + y) / 2)
}
