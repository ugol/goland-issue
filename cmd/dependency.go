package cmd

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello from another directory, %v!", name)
}
