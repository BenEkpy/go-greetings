package greetings

import "fmt"

func Greetings(pseudo string) string {
	r := fmt.Sprintf("Hello, %v ! Welcome !!", pseudo)
	return r
}