package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestImportPriceList(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestImportPriceList")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpfile.Name())

	db, err := NewSQLite3Client(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if _, err := db.db.Exec(`create table aprice_price_list ("SKU","OfferTermCode","RateCode","TermType");`); err != nil {
		t.Fatal(err)
	}

	table := "aprice_price_list"
	header := []string{
		"SKU", "OfferTermCode", "RateCode", "TermType",
	}
	records := [][]string{
		[]string{"YQHNG5NBWUE3D67S", "4NA7Y494T4", "YQHNG5NBWUE3D67S.4NA7Y494T4.6YS6EN2CT7", "Reserved"},
		[]string{"EYDF9FPAH9XHZRBR", "Z2E3P23VKM", "EYDF9FPAH9XHZRBR.Z2E3P23VKM.6YS6EN2CT7", "Reserved"},
		[]string{"WQ8JS87GX5QJ6VS4", "JRTCKXETXF", "WQ8JS87GX5QJ6VS4.JRTCKXETXF.6YS6EN2CT7", "OnDemand"},
		[]string{"74DZW9N4CZYAFK93", "HU7G6KETJZ", "74DZW9N4CZYAFK93.HU7G6KETJZ.6YS6EN2CT7", "Reserved"},
	}

	if err := db.ImportPriceList(table, header, records); err != nil {
		t.Errorf("got error: %s", err)
	}

	var count int
	expected := 4

	if err := db.db.QueryRow("select count(*) from aprice_price_list;").Scan(&count); err != nil {
		t.Error(err)
	}

	if count != expected {
		t.Errorf("wrong row count, expected: %d, got: %d", expected, count)
	}
}

func TestTableExists(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestTableExists")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpfile.Name())

	db, err := NewSQLite3Client(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	table := "aprice_price_list"

	exists1, err := db.TableExists(table)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if exists1 {
		t.Errorf("table %s should not exist", table)
	}

	if _, err := db.db.Exec(fmt.Sprintf(`create table %s ("SKU","OfferTermCode","RateCode","TermType");`, table)); err != nil {
		t.Error(err)
	}

	exists2, err := db.TableExists(table)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if exists2 {
		t.Errorf("table %s should exist", table)
	}
}
