package demo

import (
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	dfs(t, ch)
	close(ch)
}

func dfs(t *tree.Tree, ch chan int) {
	if t != nil {
		dfs(t.Left, ch)
		ch <- t.Value
		dfs(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2

		if !ok1 && !ok2 {
			break
		}
		if !ok1 || !ok2 || val1 != val2 {
			return false
		}
	}
	return true
}
