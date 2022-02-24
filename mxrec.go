//Looking up MX records (DNS records that point to mail server) for given host name//

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
	fmt.Println("Looking up MX records for given domain:" + arg)

	mxr, err := net.LookupMX(arg)
	if err != nil {
		log.Fatal(err)
	}
	for _, mx := range mxr {
		fmt.Println("Host:  \t \t Preference: \n", mx.Host, mx.Pref)
	}
}
