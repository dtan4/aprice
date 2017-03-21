package main

import (
	"fmt"
	"os"

	"github.com/dtan4/aprice/aws"
	"github.com/dtan4/aprice/aws/ec2"
)

func main() {
	csv, err := aws.RetrievePriceListCSV(aws.EC2Service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	header, records, err := ec2.ParseEC2PriceListCSV(csv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%q\n", header)
	fmt.Printf("%q\n", records)
}
