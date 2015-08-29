package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
)

const (
	ARP_FILTER = "arp"
)

func main() {
	log.Println("Starting up...")
	devs, err := pcap.FindAllDevs()
	if err != nil || len(devs) == 0 {
		log.Fatalf("No devices found, you must run this as root on OSX - try:\n sudo GOPATH=/Users/slowteetoe/go go run dash.go\n (error was: %s)\n", err)
	}
	for d := range devs {
		log.Println(devs[d].Name)
	}
	h, err := pcap.OpenLive("en0", 65536, true, pcap.BlockForever)
	if err != nil || h == nil {
		log.Fatalf("Error obtaining handle: %s\n", err)
	}
	defer h.Close()

	err = h.SetBPFFilter(ARP_FILTER)
	if err != nil {
		log.Fatalf("Unable to set filter! %s\n", err)
	}

	log.Println("Listening for ARP packets on en0")

	packetSource := gopacket.NewPacketSource(h, h.LinkType())
	for packet := range packetSource.Packets() {
		log.Println(packet) // Do something with a packet here.
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			log.Println("Ethernet layer detected.")
			ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
			log.Println("Source MAC: ", ethernetPacket.SrcMAC)
			log.Println("Destination MAC: ", ethernetPacket.DstMAC)
			// Ethernet type is typically IPv4 but could be ARP or other
			log.Println("Ethernet type: ", ethernetPacket.EthernetType)
			log.Println()
		}
	}
	log.Println("done.")
}
