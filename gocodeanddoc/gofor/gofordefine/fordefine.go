package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	//for语句可以一个条件有没有,也可以没有任何一个,循环判断三个条件语句外不能带()
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	//打开文件
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	//返回一个file的scanner，默认的分割函数是ScanLines
	scanner := bufio.NewScanner(file)
	//省略起始和递增条件,go语言无while,因此for语言只有跳出条件时
	// 直到EOF
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func forever() {
	//没有while
	//省略跳出条件
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(
		convertToBin(5),
		//convertToBin(13)此种写法错误
		//括号里如果要换行,加,号
		convertToBin(13),
	)
	printFile("abc.txt")
}
