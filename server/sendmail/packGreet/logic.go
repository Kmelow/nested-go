package packGreet

import "fmt"

// Greet returns greeting
func Greet(s string) string {
	return fmt.Sprintf("Hello %s", s)
}
