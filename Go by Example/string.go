package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contanins:", s.Contains("test", "es"))
	p("ToLower:", s.ToLower("TEST"))
	// 返回 bool 类型
	p("HasPrefix:", s.HasPrefix("test", "te"))

	// 返回索引位置
	p("Index:", s.Index("test", "t"))

	// 重复 5 次打印出来
	p("Repate:", s.Repeat("a", 5))

	//
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))

	// 去掉 -
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
}
