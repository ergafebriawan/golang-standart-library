package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.April, 16, 0, 0, 4, 0, time.UTC)
	fmt.Println(utc)

	formatter := "2006-02-12 11:04:04"
	val := "2020-10-09 10:07:04"

	valueTime, err := time.Parse(formatter, val)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}
}
