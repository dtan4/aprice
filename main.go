package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dtan4/aprice/aws"
	"github.com/dtan4/aprice/aws/ec2"
	"github.com/dtan4/aprice/database"
)

func main() {
	csv, err := aws.RetrievePriceListCSV(aws.EC2Service, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	header, records, err := ec2.ParseEC2PriceListCSV(csv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	dbPath := filepath.Join(os.Getenv("HOME"), "ec2.db")
	db, err := database.New(dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	fmt.Println(header)
	fmt.Println(records)
}
