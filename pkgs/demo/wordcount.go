package demo

import (
	"strings"
)

// 练习3：字符串单词统计
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wc := make(map[string]int)
	for _, ch := range words {
		if _, ok := wc[ch]; !ok {
			wc[ch] = 0
		}
		wc[ch] += 1
	}
	return wc
}
