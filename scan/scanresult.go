package scan

import (
	"Scanner/port"
	"fmt"
	"net"
)

type ScanResult struct {
	ScanEntry   *ScanEntry
	Source      net.IP
	Destination net.IP
	Port        *port.Port
	IsPassed    bool
}

func (s ScanResult) String() string {
	result := fmt.Sprintf("Source: %v\n", s.Source) +
		fmt.Sprintf("Destination: %v\n", s.Destination) +
		fmt.Sprintf("Port: %v\n", s.Port) + fmt.Sprintf("Is passed: %v\n", s.IsPassed) +
		fmt.Sprintf("Comment: %v\n\n\n", s.ScanEntry.Comment)
	return result
}
