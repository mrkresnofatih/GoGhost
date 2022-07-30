package cider

import (
	"testing"
)

func TestCreateNewCider_FromStringShouldHaveNonEmptyStringAsRawStringProperty(t *testing.T) {
	createNewCiderStringInput := "10.192.100.0/24"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.RawString == "" {
		t.Errorf("CreateNewCiderFromString Failed: CreateNewCider(%s) did not initialize CiderObject RawString property!", createNewCiderStringInput)
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveRawStringValueSameAsInput(t *testing.T) {
	createNewCiderStringInput := "10.192.100.0/24"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.RawString != createNewCiderStringInput {
		t.Errorf("CreateNewCiderFromString Failed: CreateNewCider(%s) did not initialize the correct CiderObject RawString property!", createNewCiderStringInput)
		return
	}
}

func TestCreateNewCider_FromStringShouldRejectCiderStringWithIncorrectFirstSegmentRange(t *testing.T) {
	createNewCiderStringInput := "256.192.100.0/24"
	_, err := CreateNewCider(createNewCiderStringInput)

	if err == nil {
		t.Errorf("CreateNewCiderFromString Failed: CreateNewCider(%s) did not catch invalid first segment range exception", createNewCiderStringInput)
		return
	}
}

func TestCreateNewCider_FromStringShouldRejectCiderStringWithIncorrectSecondSegmentRange(t *testing.T) {
	createNewCiderStringInput := "10.256.100.0/24"
	_, err := CreateNewCider(createNewCiderStringInput)

	if err == nil {
		t.Errorf("CreateNewCiderFromString Failed: CreateNewCider(%s) did not catch invalid 2nd segment range exception", createNewCiderStringInput)
		return
	}
}

func TestCreateNewCider_FromStringShouldRejectCiderStringWithIncorrectThirdSegmentRange(t *testing.T) {
	createNewCiderStringInput := "10.192.256.0/24"
	_, err := CreateNewCider(createNewCiderStringInput)

	if err == nil {
		t.Errorf("CreateNewCiderFromString Failed: CreateNewCider(%s) did not catch invalid 3rd segment range exception", createNewCiderStringInput)
		return
	}
}

func TestCreateNewCider_FromStringShouldRejectCiderStringWithIncorrectFourthSegmentRange(t *testing.T) {
	createNewCiderStringInput := "10.192.100.256/24"
	_, err := CreateNewCider(createNewCiderStringInput)

	if err == nil {
		t.Errorf("CreateNewCiderFromString Failed: CreateNewCider(%s) did not catch invalid 4th segment range exception", createNewCiderStringInput)
		return
	}
}

func TestCreateNewCider_FromStringShouldRejectCiderStringWithIncorrectFifthSegmentRange(t *testing.T) {
	createNewCiderStringInput := "10.192.100.0/34"
	_, err := CreateNewCider(createNewCiderStringInput)

	if err == nil {
		t.Errorf("CreateNewCiderFromString Failed: CreateNewCider(%s) did not catch invalid 5th segment range exception", createNewCiderStringInput)
		return
	}
}

func TestCreateNewCider_FromStringShouldWorkWithCorrectSegmentRanges(t *testing.T) {
	createNewCiderStringInput := "10.192.120.0/24"
	_, err := CreateNewCider(createNewCiderStringInput)

	if err != nil {
		t.Errorf("CreateNewCiderFromString Failed: CreateNewCider(%s) somehow caught an exception", createNewCiderStringInput)
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveSizePropertyNonDefault(t *testing.T) {
	createNewCiderStringInput := "10.192.120.0/24"

	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.Size == 0 {
		t.Errorf("CreateNewCiderFromString Failed: CreateNewCider(%s) resulted Size property with default value", createNewCiderStringInput)
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveSubnetSizePropertyNonDefault(t *testing.T) {
	createNewCiderStringInput := "10.192.120.10/24"

	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize == 0 {
		t.Errorf("CreateNewCiderFromString Failed: CreateNewCider(%s) resulted SubnetSize property with default value", createNewCiderStringInput)
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize1(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/1"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(2147483648) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize4(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/4"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(268435456) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize7(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/7"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(33554432) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize10(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/10"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(4194304) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize13(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/13"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(524288) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize16(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/16"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(65536) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize19(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/19"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(8192) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize22(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/22"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(1024) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize25(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/25"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(128) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize28(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/28"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(16) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldHaveCorrectSubnetSizeForSize31(t *testing.T) {
	createNewCiderStringInput := "0.0.0.0/31"
	ciderObject, err := CreateNewCider(createNewCiderStringInput)
	handleErrorCausedByBadTestCase(t, err)

	if ciderObject.SubnetSize != uint32(2) {
		t.Errorf("CreateNewCiderFromString Failed: Calculate Subnet result is incorrect!")
		return
	}
}

func TestCreateNewCider_FromStringShouldRejectIncorrectNumberOfSegmentsInInput(t *testing.T) {
	createNewCiderStringInput := "0.0.0/16"
	_, err := CreateNewCider(createNewCiderStringInput)

	if err == nil {
		t.Errorf("Incorrect Number of segments uncaught!")
		return
	}
}

func TestCreateNewCidr_FromStringShouldCatchCidrSizeAndCidrAddressMismatch(t *testing.T) {
	createNewCiderStringInput := "10.192.128.0/16"
	_, err := CreateNewCider(createNewCiderStringInput)

	if err == nil {
		t.Errorf("CidrSize & CidrAddress Mismatch is uncaught!")
		return
	}
}

func handleErrorCausedByBadTestCase(t *testing.T, err error) {
	if err != nil {
		t.Errorf("BadTestCase: " + err.Error())
	}
}
