package cider

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type Cider struct {
	RawString  string
	Size       int
	SubnetSize uint32
}

func CreateNewCider(ciderString string) (*Cider, error) {
	_, err := validateCiderString(&ciderString)
	if err != nil {
		newError := errors.New("Invalid Cider string: " + err.Error())
		return nil, newError
	}

	cider := Cider{}
	cider.RawString = ciderString

	sizeOfCiderFromInput, err := strconv.Atoi(strings.Split(ciderString, "/")[1])
	if err != nil {
		newError := errors.New("Can't convert size to string: " + err.Error())
		return nil, newError
	}

	cider.Size = sizeOfCiderFromInput

	cider.SubnetSize = getSubnetSizeFromSize(sizeOfCiderFromInput)

	return &cider, nil
}

func validateCiderString(ciderString *string) (bool, error) {
	isFirstSegmentRangeValid, err := validateSegmentsRange(ciderString)
	if err != nil {
		return false, err
	}
	return isFirstSegmentRangeValid, nil
}

func validateSegmentsRange(ciderString *string) (bool, error) {
	ciderStringAddress := strings.Split(*ciderString, "/")[0]
	segments := strings.Split(ciderStringAddress, ".")
	if len(segments) != 4 {
		incorrectNumberOfSegmentsInCidrAddressError := errors.New("Incorrect Number of segments")
		return false, incorrectNumberOfSegmentsInCidrAddressError
	}
	for index, value := range []int{1, 2, 3, 4, 5} {
		if value == 5 {
			ciderSizeSegment := strings.Split(*ciderString, "/")[1]
			currentSegment, err := strconv.Atoi(ciderSizeSegment)
			if err != nil {
				stringToIntConversionError := errors.New("cannot convert string to Int")
				return false, stringToIntConversionError
			}

			if currentSegment < 1 || currentSegment > 32 {
				invalidSegmentError := errors.New("current segment is < 0 or > 255")
				return false, invalidSegmentError
			}
			continue
		}

		currentSegment, err := strconv.Atoi(segments[index])
		if err != nil {
			stringToIntConversionError := errors.New("cannot convert string to Int")
			return false, stringToIntConversionError
		}

		if currentSegment < 0 || currentSegment > 255 {
			invalidSegmentError := errors.New("current segment is < 0 or > 255")
			return false, invalidSegmentError
		}
	}
	return true, nil
}

func getSubnetSizeFromSize(size int) uint32 {
	return uint32(math.Pow(2, float64(32-size)))
}
