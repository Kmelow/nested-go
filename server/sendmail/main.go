package main

import (
	"fmt"

	"github.com/kmelow/nested-go/server/sendmail/packGreet"
)

func main() {
	fmt.Println(packGreet.Greet("Shoshanna"))
}
