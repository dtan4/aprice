package aws

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// RetrievePriceListCSV downloads price list CSV
func RetrievePriceListCSV(service string) (string, error) {
	url := priceListCSVURL(service)

	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve price list CSV")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read price list CSV body")
	}

	return string(body), nil
}

func priceListCSVURL(service string) string {
	return fmt.Sprintf("https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/%s/current/index.csv", service)
}
