package demo

import (
	"io"
	"os"
	"strings"
)

// io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。
type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = transform(b[i])
	}
	return n, err
}

// 向后映射13位
func transform(ch byte) byte {
	if 'a' <= ch && ch <= 'z' {
		return (ch-'a'+13)%26 + 'a'
	} else if 'A' <= ch && ch <= 'Z' {
		return (ch-'A'+13)%26 + 'A'
	} else {
		return ch
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
