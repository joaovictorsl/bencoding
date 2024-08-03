package bencoding

import (
	"fmt"
	"reflect"
)

var (
	ErrMinusZeroInteger = fmt.Errorf("-0 is not a valid integer")
)

type ErrDataCastFail struct {
	castedData interface{}
	data       interface{}
}

func NewErrCastFail(data interface{}, castedData interface{}) ErrDataCastFail {
	return ErrDataCastFail{
		castedData: castedData,
		data:       data,
	}
}

func (err ErrDataCastFail) Error() string {
	return fmt.Sprintf("failed to cast data to %v, data is of type %v", reflect.TypeOf(err.castedData), reflect.TypeOf(err.data))
}

type ErrLeadingZeroInteger struct {
	invalidInt string
}

func NewErrLeadingZeroInteger(invalidInt string) ErrLeadingZeroInteger {
	return ErrLeadingZeroInteger{
		invalidInt: invalidInt,
	}
}

func (err ErrLeadingZeroInteger) Error() string {
	return fmt.Sprintf("%s is not a valid integer value, leading 0 are not allowed", err.invalidInt)
}

type ErrInvalidType struct {
	invalidType reflect.Type
}

func NewErrInvalidType(invalidType interface{}) ErrInvalidType {
	return ErrInvalidType{
		invalidType: reflect.TypeOf(invalidType),
	}
}

func (err ErrInvalidType) Error() string {
	return fmt.Sprintf("value of type %v cannot be encoded, only int, string, []interface{} and map[string]interface{} can be encoded", err.invalidType)
}
