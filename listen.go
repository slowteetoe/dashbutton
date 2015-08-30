package main

import (
	"bytes"
	"flag"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
)

func gatoradeButtonPress() {
	log.Println("Pressed the Gatorade button.")
}

func gladButtonPress() {
	log.Println("Pressed the Glad button.")
}

var interfaceName = flag.String("inf", "en0", "Provide a valid interface name to listen on (e.g. eth0 or en0)")

func main() {
	gatoradeDashButton, _ := net.ParseMAC("74:75:48:a4:59:a8")
	gladDashButton, _ := net.ParseMAC("74:75:48:29:a8:7c")

	flag.Parse()
	log.Printf("Starting up on interface[%v]...", *interfaceName)

	h, err := pcap.OpenLive(*interfaceName, 65536, true, pcap.BlockForever)

	if err != nil || h == nil {
		log.Fatalf("Error opening interface: %s\nPerhaps you need to run as root?\n", err)
	}
	defer h.Close()

	err = h.SetBPFFilter("arp and ((ether src host " + gatoradeDashButton.String() + ") or (ether src host " + gladDashButton.String() + "))")
	if err != nil {
		log.Fatalf("Unable to set filter! %s\n", err)
	}
	log.Println("Listening for Dash buttons...")

	packetSource := gopacket.NewPacketSource(h, h.LinkType())

	// Since we're using a BPF filter to limit packets to only our buttons, we don't need to worry about anything besides MAC here...
	for packet := range packetSource.Packets() {
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		if bytes.Equal(ethernetPacket.SrcMAC, gatoradeDashButton) {
			gatoradeButtonPress()
		} else if bytes.Equal(ethernetPacket.SrcMAC, gladDashButton) {
			gladButtonPress()
		} else {
			log.Printf("Received button press, but don't know how to handle MAC[%v]", ethernetPacket.SrcMAC)
		}
	}
}
