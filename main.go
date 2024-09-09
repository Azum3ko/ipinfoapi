package main

import (
	"fmt"
)


func main() {
	ipinfo := Ipinfo{
		ip: "170.246.146.161",
	}
	fmt.Println(ipinfo.Getall())
}