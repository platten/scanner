package scan

import (
	"Scanner/port"
	"fmt"
	"log"
	"net"
	"time"
)

func getMyIPs() net.IP {
	conn, err := net.Dial("tcp", "1.1.1.1:443")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.TCPAddr).IP
}

// Func for checking if current IP or NAT IP is in list of sources.
func isSource(sources []net.IPNet) bool {
	ip := getMyIPs()
	for _, source := range sources {
		if source.Contains(ip) {
			return true
		}
	}
	return false
}

func resolveHostname(hostname string) []net.IP {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		log.Fatal(err)
	}
	return ips
}

func CheckIfPortOpen(destination net.IP, port *port.Port) bool {
	portString := fmt.Sprintf("%d", port.Port)
	destinationString := fmt.Sprintf("%s", destination)
	destinationPortString := net.JoinHostPort(destinationString, portString)
	switch port.Protocol {
	case TCP:
		conn, err := net.DialTimeout("tcp", destinationPortString, time.Second)
		if err != nil {
			fmt.Println("Connecting error:", err)
			return false
		}
		fmt.Println("Successfully Opened", destinationPortString)
		conn.Close()
		return true
	case UDP:
		conn, err := net.DialTimeout("udp", destinationPortString, time.Second)
		if err != nil {
			fmt.Println("Connecting error:", err)
			return false
		}
		fmt.Println("Successfully Opened", destinationPortString)
		conn.Close()
		return true
	case ICMP:
		conn, err := net.DialTimeout("icmp", destinationString, time.Second)
		if err != nil {
			fmt.Println("Connecting error:", err)
			return false
		}
		fmt.Println("Successfully Opened", destinationPortString)
		conn.Close()
		return true
	}
	fmt.Println("Invalid Protocol")
	return false
}
