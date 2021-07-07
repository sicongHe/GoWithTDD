package main

import (
	"fmt"

	"github.com/siconghe/hello/concurrency"
)

func mockWebsitechecker(_ string) bool {
	return true
}
func main() {
	urls := []string{
		"www.1.com",
		"www.2.com",
		"www.3.com",
	}
	got := concurrency.CheckWebsite(mockWebsitechecker, urls)
	fmt.Printf("%#v", got)
}
