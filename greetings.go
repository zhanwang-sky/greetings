package greetings

import "fmt"

func Hello(name string) string {
	msg := fmt.Sprintf("Hello %v", name)
	return msg
}
