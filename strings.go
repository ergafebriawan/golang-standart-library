package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("beli baju", "beli"))
	fmt.Println(strings.Split("erga febriawan", " "))
	fmt.Println(strings.ToLower("Erga"))
	fmt.Println(strings.ToUpper("erga"))
	fmt.Println(strings.Trim("    erga febrian     ", " "))
	fmt.Println(strings.ReplaceAll("erga febrian", "febrian", "febriawan"))
}
