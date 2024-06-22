package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "put your database")
	var port *int = flag.Int("port", 0, "put your port")
	var username *string = flag.String("username", "root", "put your username")
	var password *string = flag.String("password", "root", "put your password")

	flag.Parse()

	fmt.Println("host", *host)
	fmt.Println("port", *port)
	fmt.Println("username", *username)
	fmt.Println("password", *password)

	//contoh untuk menjalankan
	//go run flag.go -host="11.11.1.5" -port=3306 -username=erga -password=rahasia
}
