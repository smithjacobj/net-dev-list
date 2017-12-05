package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	devs, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}
	for _, dev := range devs {
		fmt.Println(dev)
	}
}
