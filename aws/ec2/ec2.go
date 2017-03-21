package ec2

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/pkg/errors"
)

const (
	metadataLineCount = 5
)

// ParseEC2PriceListCSV converts EC2 price list CSV to Go object
func ParseEC2PriceListCSV(body string) ([]string, [][]string, error) {
	lines := strings.Split(body, "\n")
	bodyWithoutMetadata := strings.Join(lines[metadataLineCount:], "\n")
	reader := csv.NewReader(strings.NewReader(bodyWithoutMetadata))

	header, err := reader.Read()
	if err != nil {
		return []string{}, [][]string{}, errors.Wrap(err, "failed to parse header line")
	}

	records := [][]string{}

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return []string{}, [][]string{}, errors.Wrap(err, "failed to parse line")
		}

		records = append(records, record)
	}

	return header, records, nil
}
