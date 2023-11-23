/*
Copyright © 2023 Paul Pietkiewicz <pawel.pietkiewicz@gmail.com>
*/
package main

import (
	"Scanner/scan"
	"fmt"
)

func main() {
	s := scan.NewScanEntry("scanme.nmap.org", nil, nil, "TCP:443,TCP:80", "Test")
	fmt.Println(s)
	s.Scan(false)
	results := s.Scan(false)
	for _, result := range results {
		fmt.Print(result)
	}
	allresults, err := s.ScansPassed()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("passed: " + fmt.Sprintf("%v", allresults))
}
