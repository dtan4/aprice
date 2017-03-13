package aws

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// RetrievePriceList downloads price list JSON
func RetrievePriceList(service string) (string, error) {
	url := priceListURL(service)

	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve price list JSON")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read price list JSON body")
	}

	return string(body), nil
}

func priceListURL(service string) string {
	return fmt.Sprintf("https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/%s/current/index.json", service)
}
