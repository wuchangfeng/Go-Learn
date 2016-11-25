package main

// 整个文件读取到一个字符串里面

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "A.txt"
	outputFile := "B.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	// WriteFile 将 []byte 写到 文件中
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
