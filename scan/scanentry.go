package scan

import (
	"Scanner/port"
	"errors"
	"fmt"
	"net"
)

const (
	TCP  int = 0
	UDP      = 1
	ICMP     = 2
)

type ScanEntry struct {
	Source              []net.IPNet
	SourceApplicable    bool
	DestinationHostname string
	Destination         []net.IP
	Ports               []*port.Port
	IsScanned           bool
	Results             []*ScanResult
	Comment             string
}

func NewScanEntry(destinationHostname string, sources []net.IPNet, destinations []net.IP, ports string, comment string) *ScanEntry {
	p := new(ScanEntry)
	p.Source = sources
	p.DestinationHostname = destinationHostname
	if destinations == nil {
		p.Destination = resolveHostname(destinationHostname)
	} else {
		p.Destination = destinations
	}
	p.Ports = port.NewPorts(ports)
	p.IsScanned = false
	p.Results = nil
	p.SourceApplicable = isSource(sources)
	p.Comment = comment
	return p
}

func (s *ScanEntry) Scan(filtered bool) []*ScanResult {
	var results []*ScanResult
	if (filtered && s.SourceApplicable) || !filtered {
		for _, destination := range s.Destination {
			for _, port := range s.Ports {
				result := new(ScanResult)
				result.ScanEntry = s
				result.Source = getMyIPs()
				result.Destination = destination
				result.Port = port
				result.IsPassed = CheckIfPortOpen(destination, port)
				results = append(results, result)
				s.Results = append(s.Results, result)
			}
		}
	}
	return results
}

func (s *ScanEntry) ScansPassed() (bool, error) {
	if len(s.Results) == 0 {
		return false, errors.New("Scan not performed")
	}
	for _, scanResult := range s.Results {
		if scanResult.IsPassed {
			return true, nil
		}
	}
	return false, nil
}

func (s ScanEntry) String() string {
	result := fmt.Sprintf("Source: %v\n", s.Source) + fmt.Sprintf("Destination: %v\n", s.Destination) + fmt.Sprintf("Port: %v\n", s.Ports) + fmt.Sprintf("Comment: %v\n", s.Comment)
	return result
}

func (s *ScanEntry) GetTCPPorts() []int {
	var tcpPorts []int
	for _, port := range s.Ports {
		if port.Protocol == TCP {
			tcpPorts = append(tcpPorts, port.Port)
		}
	}
	return tcpPorts
}

func (s *ScanEntry) GetUDPPorts() []int {
	var udpPorts []int
	for _, port := range s.Ports {
		if port.Protocol == UDP {
			udpPorts = append(udpPorts, port.Port)
		}
	}
	return udpPorts
}
