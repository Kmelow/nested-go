package main

import (
	"fmt"

	"github.com/kmelow/nested-go/server/sendmail/packgreet"
)

func main() {
	fmt.Println("Changing the sendmail service")
	fmt.Println(packgreet.Greet("Shoshanna"))
}
