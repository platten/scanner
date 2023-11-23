package port

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	TCP  int = 0
	UDP      = 1
	ICMP     = 2
)

type Port struct {
	Protocol int
	Port     int
}

func (p Port) String() string {
	switch p.Protocol {
	case TCP:
		return fmt.Sprintf("TCP:%d", p.Port)
	case UDP:
		return fmt.Sprintf("TCP:%d", p.Port)
	case ICMP:
		return "ICMP"
	}
	return "unknown"
}

func NewPorts(portstring string) []*Port {
	var ports []*Port
	for _, portstring := range strings.Split(portstring, ",") {
		p := new(Port)
		portstringlist := strings.Split(portstring, ":")
		switch portstringlist[0] {
		case "TCP":
			p.Protocol = TCP
			p.Port, _ = strconv.Atoi(string(portstringlist[1]))
		case "UDP":
			p.Protocol = UDP
			p.Port, _ = strconv.Atoi(string(portstringlist[1]))
		case "ICMP":
			p.Protocol = ICMP
			p.Port = 0
		}
		ports = append(ports, p)
	}
	return ports
}

func (p Port) CheckIfPortOpen(destination net.IP) bool {
	portString := fmt.Sprintf("%d", p.Port)
	destinationString := fmt.Sprintf("%s", destination)
	destinationPortString := net.JoinHostPort(destinationString, portString)
	switch p.Protocol {
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
