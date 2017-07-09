package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dtan4/aprice/aws"
	"github.com/dtan4/aprice/aws/ec2"
	"github.com/dtan4/aprice/cmd/util"
	"github.com/dtan4/aprice/db"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	dbname = "ec2.db"
	table  = "aprice_price_list"
)

// ec2InitCmd represents the ec2 init command
var ec2InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize for EC2 price list",
	RunE:  doEC2Init,
}

func doEC2Init(cmd *cobra.Command, args []string) error {
	dirExists, err := util.DirExists(apriceDir)
	if err != nil {
		return errors.Wrapf(err, "failed to check whether %q exists or not", apriceDir)
	}

	if !dirExists {
		fmt.Printf("===> creating aprice data directory %q...\n", apriceDir)

		if err := os.MkdirAll(apriceDir, 0755); err != nil {
			return errors.Wrapf(err, "failed to create %q", apriceDir)
		}
	}

	fmt.Println("===> retrieving price list...")

	csv, err := aws.RetrievePriceListCSV(aws.EC2Service, true)
	if err != nil {
		return errors.Wrap(err, "failed to retreive EC2 price list")
	}

	fmt.Println("===> parsing price list...")

	header, records, err := ec2.ParseEC2PriceListCSV(csv)
	if err != nil {
		return errors.Wrap(err, "failed to parse EC2 price list CSV")
	}

	d, err := db.NewSQLite3Client(filepath.Join(apriceDir, dbname))
	if err != nil {
		return errors.Wrap(err, "failed to initialize SQLite3 client")
	}

	fmt.Println("===> opening table...")

	exists, err := d.TableExists(table)
	if err != nil {
		return errors.Wrapf(err, "failed to check whether table %q exists or not", table)
	}

	if !exists {
		if err := d.CreateTable(table, header); err != nil {
			return errors.Wrapf(err, "failed to create table %q", table)
		}
	}

	fmt.Println("===> inserting records...")

	if err := d.ImportPriceList(table, header, records); err != nil {
		return errors.Wrap(err, "failed to import price list to database")
	}

	fmt.Println("complete!")

	return nil
}

func init() {
	ec2Cmd.AddCommand(ec2InitCmd)
}
