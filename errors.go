package bencoding

import "fmt"

var (
	ErrMinusZeroInteger   = fmt.Errorf("-0 is not a valid integer")
	ErrLeadingZeroInteger = fmt.Errorf("an integer with a leading 0 is not valid")
)
