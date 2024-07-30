package greetings

import "fmt"

func Greetings(pseudo string) string {
	r := fmt.Sprintf("Hello, Welcome %v !!", pseudo)
	return r
}