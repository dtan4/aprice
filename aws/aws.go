package aws

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/cheggaaa/pb.v1"
)

const (
	// EC2Service represents EC2 service name
	EC2Service = "AmazonEC2"
)

// RetrievePriceListCSV downloads price list CSV
func RetrievePriceListCSV(service string, showProgress bool) (string, error) {
	url := priceListCSVURL(service)

	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve price list CSV")
	}
	defer resp.Body.Close()

	i, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return "", errors.Wrap(err, "failed to get Content-Length")
	}

	var reader io.Reader

	if showProgress {
		bar := pb.New(i).SetUnits(pb.U_BYTES).SetRefreshRate(10 * time.Millisecond)
		bar.ShowSpeed = true
		bar.Start()
		defer bar.Finish()

		reader = bar.NewProxyReader(resp.Body)
	} else {
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", errors.Wrap(err, "failed to read price list CSV body")
	}

	return string(body), nil
}

func priceListCSVURL(service string) string {
	return fmt.Sprintf("https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/%s/current/index.csv", service)
}
