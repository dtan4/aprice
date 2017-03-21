package aws

import (
	"testing"

	"github.com/h2non/gock"
)

func TestRetrievePriceListCSV(t *testing.T) {
	defer gock.Off()

	gock.New("https://pricing.us-east-1.amazonaws.com").
		Get("/offers/v1.0/aws/AmazonEC2/current/index.csv").
		Reply(200).
		BodyString(`"FormatVersion","v1.0"
"Disclaimer","This pricing list is for informational purposes only. All prices are subject to the additional terms included in the pricing pages on http://aws.amazon.com. All Free Tier prices are also subject to the terms included at https://aws.amazon.com/free/"
"Publication Date","2017-03-02T18:32:21Z"
"Version","20170302183221"
"OfferCode","AmazonEC2"
"SKU","OfferTermCode","RateCode","TermType","PriceDescription","EffectiveDate","StartingRange","EndingRange","Unit","PricePerUnit","Currency","LeaseContractLength","PurchaseOption","OfferingClass","Product Family","serviceCode","Location","Location Type","Instance Type","Current Generation","Instance Family","vCPU","Physical Processor","Clock Speed","Memory","Storage","Network Performance","Processor Architecture","Storage Media","Volume Type","Max Volume Size","Max IOPS/volume","Max IOPS Burst Performance","Max throughput/volume","Provisioned","Tenancy","EBS Optimized","Operating System","License Model","Group","Group Description","Transfer Type","From Location","From Location Type","To Location","To Location Type","usageType","operation","Comments","Dedicated EBS Throughput","Enhanced Networking Supported","GPU","Instance Capacity - 10xlarge","Instance Capacity - 2xlarge","Instance Capacity - 4xlarge","Instance Capacity - 8xlarge","Instance Capacity - large","Instance Capacity - medium","Instance Capacity - xlarge","Intel AVX Available","Intel AVX2 Available","Intel Turbo Available","Physical Cores","Pre Installed S/W","Processor Features","Sockets"
"YQHNG5NBWUE3D67S","4NA7Y494T4","YQHNG5NBWUE3D67S.4NA7Y494T4.6YS6EN2CT7","Reserved","Red Hat Enterprise Linux (Amazon VPC), m4.xlarge instance-hours used this month","2016-11-30","0","Inf","Hrs","0.2230000000","USD","1yr","No Upfront","standard","Compute Instance","AmazonEC2","US East (Ohio)","AWS Region","m4.xlarge","Yes","General purpose","4","Intel Xeon E5-2676 v3 (Haswell)","2.4  GHz","16 GiB","EBS only","High","64-bit",,,,,,,,"Dedicated",,"RHEL","No License required",,,,,,,,"USE2-DedicatedUsage:m4.xlarge","RunInstances:0010",,"750 Mbps","Yes",,,,,,,,,,,,,"NA","Intel AVX; Intel AVX2; Intel Turbo",
`).AddHeader("Content-Length", "100")

	service := "AmazonEC2"

	got, err := RetrievePriceListCSV(service, false)
	if err != nil {
		t.Errorf("error should not be raised: %s", err)
	}

	expected := `"FormatVersion","v1.0"
"Disclaimer","This pricing list is for informational purposes only. All prices are subject to the additional terms included in the pricing pages on http://aws.amazon.com. All Free Tier prices are also subject to the terms included at https://aws.amazon.com/free/"
"Publication Date","2017-03-02T18:32:21Z"
"Version","20170302183221"
"OfferCode","AmazonEC2"
"SKU","OfferTermCode","RateCode","TermType","PriceDescription","EffectiveDate","StartingRange","EndingRange","Unit","PricePerUnit","Currency","LeaseContractLength","PurchaseOption","OfferingClass","Product Family","serviceCode","Location","Location Type","Instance Type","Current Generation","Instance Family","vCPU","Physical Processor","Clock Speed","Memory","Storage","Network Performance","Processor Architecture","Storage Media","Volume Type","Max Volume Size","Max IOPS/volume","Max IOPS Burst Performance","Max throughput/volume","Provisioned","Tenancy","EBS Optimized","Operating System","License Model","Group","Group Description","Transfer Type","From Location","From Location Type","To Location","To Location Type","usageType","operation","Comments","Dedicated EBS Throughput","Enhanced Networking Supported","GPU","Instance Capacity - 10xlarge","Instance Capacity - 2xlarge","Instance Capacity - 4xlarge","Instance Capacity - 8xlarge","Instance Capacity - large","Instance Capacity - medium","Instance Capacity - xlarge","Intel AVX Available","Intel AVX2 Available","Intel Turbo Available","Physical Cores","Pre Installed S/W","Processor Features","Sockets"
"YQHNG5NBWUE3D67S","4NA7Y494T4","YQHNG5NBWUE3D67S.4NA7Y494T4.6YS6EN2CT7","Reserved","Red Hat Enterprise Linux (Amazon VPC), m4.xlarge instance-hours used this month","2016-11-30","0","Inf","Hrs","0.2230000000","USD","1yr","No Upfront","standard","Compute Instance","AmazonEC2","US East (Ohio)","AWS Region","m4.xlarge","Yes","General purpose","4","Intel Xeon E5-2676 v3 (Haswell)","2.4  GHz","16 GiB","EBS only","High","64-bit",,,,,,,,"Dedicated",,"RHEL","No License required",,,,,,,,"USE2-DedicatedUsage:m4.xlarge","RunInstances:0010",,"750 Mbps","Yes",,,,,,,,,,,,,"NA","Intel AVX; Intel AVX2; Intel Turbo",
`
	if got != expected {
		t.Errorf("body does not match. expected: %q, got: %q", expected, got)
	}
}
