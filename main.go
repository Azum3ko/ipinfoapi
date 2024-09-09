package main

import (
	"fmt"
)


func main() {
	ipinfo := Ipinfo{
		ip: "170.246.146.161",
		token: "3eb99b5b4d2884",
	}
	fmt.Println(ipinfo.City())

}