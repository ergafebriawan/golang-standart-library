package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))

	fmt.Println(regex.FindAllString("eko edo e1o eeo", 10))
}
