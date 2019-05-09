package main

import "fmt"

func hello(name string) string {
	return fmt.Sprintf("hello unexported from same directory, %v!", name)
}

func Hello(name string) string {
	return fmt.Sprintf("hello exported from same directory, %v!", name)
}
