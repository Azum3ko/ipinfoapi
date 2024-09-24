package main

import (
	"fmt"
	"pkg/ip.go"
)


func main() {
	ipinfo := Ipinfo{
		ip: "",
		token: "3eb99b5b4d2884",
	}
	fmt.Println(ipinfo.City())

}