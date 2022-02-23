//Looking up host names for given IP address

package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No valid IP Address provided")
	}
	arg := os.Args[1]
	ip := net.ParseIP(arg) //Divide into octets for validation//
	if ip == nil {
		log.Fatal("Given IP is not valid")
	}
	fmt.Println("Looking up hostnames for IP address:" + arg)
	hosts, err := net.LookupAddr(ip.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, host := range hosts {
		fmt.Println(host)
	}
}
