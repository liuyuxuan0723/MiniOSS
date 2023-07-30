package demo

// 练习7: 实现一个 Reader 类型，产生一个 ASCII 字符 'A' 的无限流。
type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}
