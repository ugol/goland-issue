package main

import (
	"fmt"
	"hello-module/cmd"
)

func main() {
	fmt.Println("Module testing outside GOPATH")
	fmt.Println(silly("Test"))
	fmt.Println(hello("World"))
	fmt.Println(Hello("World"))
	fmt.Println(cmd.Hello("Dependency"))
	fmt.Println(cmd.HelloOther("Other Dependency"))
	fmt.Println(cmd.HelloWrapper("Other Dependency (wrapped)"))
}

func silly(name string) string {
	return fmt.Sprintf("hello from same file, %v!", name)
}