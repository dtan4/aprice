package aws

import (
	"testing"

	"github.com/h2non/gock"
)

func TestRetrievePriceList(t *testing.T) {
	defer gock.Off()

	gock.New("https://pricing.us-east-1.amazonaws.com").Get("/offers/v1.0/aws/AmazonEC2/current/index.json").Reply(200).BodyString("{}")

	service := "AmazonEC2"

	got, err := RetrievePriceList(service)
	if err != nil {
		t.Errorf("error should not be raised: %s", err)
	}

	expected := "{}"
	if got != expected {
		t.Errorf("body does not match. expected: %q, got: %q", expected, got)
	}
}
