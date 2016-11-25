package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "input.txt"
	outputFile := "products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	// 例如os.Open函数在打开文件失败时将返回一个不为nil的error变量
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
