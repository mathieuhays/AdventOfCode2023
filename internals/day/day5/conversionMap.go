package day5

import (
	"errors"
)

type ConversionMap struct {
	DestinationStart int
	SourceStart      int
	Size             int
	SourceEnd        int
}

func (c *ConversionMap) InRange(seed int) bool {
	return seed >= c.SourceStart &&
		seed <= c.SourceEnd
}

func (c *ConversionMap) Convert(seed int) (int, error) {
	if !c.InRange(seed) {
		return 0, errors.New("conversion error. input out of bound")
	}

	return seed - c.SourceStart + c.DestinationStart, nil
}
