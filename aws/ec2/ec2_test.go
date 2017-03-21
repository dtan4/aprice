package ec2

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestParseEC2PriceListCSV(t *testing.T) {
	body, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "ec2.csv"))
	if err != nil {
		t.Fatalf("failed to open testdata: %s", err)
	}

	_, records, err := ParseEC2PriceListCSV(string(body))
	if err != nil {
		t.Errorf("error should not be raised: %s", err)
	}

	if len(records) != 100 {
		t.Errorf("number of records should be 100, got: %d", len(records))
	}
}
