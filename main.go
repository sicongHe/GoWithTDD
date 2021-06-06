package main

import (
	"fmt"

	"github.com/siconghe/hello/greeting"
)

func main() {
	fmt.Println(greeting.Hi("Sicong", ""))
	fmt.Println(greeting.Hello())
}
