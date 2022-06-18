package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	//Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\n Name:", device.Name)
		fmt.Println("\n Description:", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("IP Address:", address.IP)
			fmt.Println("Subnet mask:", address.Netmask)

		}
	}
}
