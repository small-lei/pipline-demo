package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, Kubernetes！I'm from Jenkins CI！")
	fmt.Println("hello, welcome zhangxiaolei image!")
	fmt.Println("来了老弟！！！！")
	fmt.Println("修改一下今天！")

	//1、一次性读取文件内容,还有一个 ReadAll的函数，也能读取
	data, err := ioutil.ReadFile("./conf/config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
