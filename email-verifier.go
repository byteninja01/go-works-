package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord")

	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" {
			continue
		}
		checkDomain(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// --- Check MX ---
	mxRecords, err := net.LookupMX(domain)
	if err == nil && len(mxRecords) > 0 {
		hasMX = true
	}

	// --- Check SPF ---
	txtRecords, err := net.LookupTXT(domain)
	if err == nil {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	}

	// --- Check DMARC ---
	dmarcTXT, err := net.LookupTXT("_dmarc." + domain)
	if err == nil {
		for _, record := range dmarcTXT {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	}

	// --- Print clean CSV line ---
	fmt.Printf("%s,%t,%t,%q,%t,%q\n",
		domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
