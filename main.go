package main

import (
	"fmt"
	"os"

	"github.com/dtan4/aprice/aws"
	"github.com/dtan4/aprice/aws/ec2"
	"github.com/dtan4/aprice/db"
)

const (
	// TODO: use absolute path
	dbFilename = "aprice.db"
	table      = "aprice_price_list"
)

func main() {
	fmt.Println("===> retrieving price list...")

	csv, err := aws.RetrievePriceListCSV(aws.EC2Service, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	fmt.Println("===> parsing price list...")

	header, records, err := ec2.ParseEC2PriceListCSV(csv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	d, err := db.NewSQLite3Client(dbFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	fmt.Println("===> opening table...")

	exists, err := d.TableExists(table)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	if !exists {
		if err := d.CreateTable(table, header); err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("===> inserting records...")

	if err := d.ImportPriceList(table, header, records); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	fmt.Println("complete!")
}
