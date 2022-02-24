//Looking up nameservers for given domain name//

package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No valid domain provided")
	}
	arg := os.Args[1]
	fmt.Println("Looking up nameservers for given domain:" + arg)

	ns, err := net.LookupNS(arg)
	if err != nil {
		log.Fatal(err)
	}
	for _, ns := range ns {
		fmt.Println(ns.Host)
	}
}
