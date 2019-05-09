package cmd

import "fmt"

func HelloOther(name string) string {
	return fmt.Sprintf("Hello from another directory, different file, %v!", name)
}

func HelloWrapper(name string) string {
	return Hello(name)
}

