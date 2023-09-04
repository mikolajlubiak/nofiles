package main

import (
	"nofiles/antiav"
	"log"
)


func main() {
	antiav.Sandbox()
	err := antiav.Before()
	if err != nil {
		log.Println("Preparation failed:", err)
	}

//	ips, err := antiav.Hosts()
//	if err != nil {
//		log.Println("Scanning failed:", err)
//	}
//	ports := antiav.Ports(ips)

	antiav.After()
	if err != nil {
		log.Println("Finishing failed:", err)
	}
}
