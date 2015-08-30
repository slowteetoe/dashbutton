package main

import (
	"bytes"
	"flag"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
)

var emptySource = []byte{0, 0, 0, 0}

var interfaceName = flag.String("inf", "en0", "Provide a valid interface name to listen on (e.g. eth0 or en0)")

func main() {
	log.Println("Starting up...")
	flag.Parse()
	devs, err := pcap.FindAllDevs()
	if err != nil || len(devs) == 0 {
		log.Printf("No devices found, you must run this as 'root' on OSX - try:\n sudo GOPATH=/Users/slowteetoe/go go run identify.go\n (error was: %s)\n", err)
	}
	h, err := pcap.OpenLive(*interfaceName, 65536, true, pcap.BlockForever)
	if err != nil || h == nil {
		log.Printf("Error obtaining handle: %s\n", err)
		log.Println("Valid interfaces on this device:")
		for d := range devs {
			log.Println(devs[d].Name)
		}
		return
	}
	defer h.Close()

	err = h.SetBPFFilter("arp")
	if err != nil {
		log.Fatalf("Unable to set filter! %s\n", err)
	}

	log.Println("Listening for ARP packets on en0 to identify dash button.\nPress the dash button and watch the screen. You will want the 'Source MAC'")

	packetSource := gopacket.NewPacketSource(h, h.LinkType())
	for packet := range packetSource.Packets() {

		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)

		arpLayer := packet.Layer(layers.LayerTypeARP)
		arp, _ := arpLayer.(*layers.ARP)

		if bytes.Equal(arp.SourceProtAddress, emptySource) {
			log.Printf("ARP packet, Source MAC[%v], Destination MAC[%v]\n", ethernetPacket.SrcMAC, ethernetPacket.DstMAC)
		}
	}

}
