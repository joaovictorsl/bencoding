package bencoding

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func DecodeTo[T any](encodedData *bufio.Reader) (castedData T, err error) {
	var zero T

	data, err := decodeNext(encodedData)
	if err != nil {
		return zero, err
	}

	castedData, ok := data.(T)
	if !ok {
		return zero, NewErrCastFail(data, castedData)
	}

	return castedData, nil
}

func Decode(encodedData *bufio.Reader) (data interface{}, err error) {
	return decodeNext(encodedData)
}

func decodeNext(encodedData *bufio.Reader) (data interface{}, err error) {
	var decode func(encodedData *bufio.Reader) (data interface{}, err error)
	b, err := encodedData.ReadByte()
	if err != nil {
		return nil, err
	}

	if unicode.IsDigit(rune(b)) {
		encodedData.UnreadByte()
		decode = decodeString
	} else if b == 'i' {
		decode = decodeInteger
	} else if b == 'l' {
		decode = decodeList
	} else if b == 'd' {
		decode = decodeDictionary
	} else {
		return nil, fmt.Errorf("error while decoding, next element is invalid")
	}

	return decode(encodedData)
}

func decodeString(encodedData *bufio.Reader) (data interface{}, err error) {
	lengthBytes := make([]byte, 0)
	for {
		b, err := encodedData.ReadByte()
		if err != nil {
			return nil, err
		} else if b == ':' {
			break
		}

		lengthBytes = append(lengthBytes, b)
	}

	length, err := strconv.Atoi(string(lengthBytes))
	if err != nil {
		return nil, err
	}

	dataBytes := make([]byte, 0)
	for ; length > 0; length-- {
		b, err := encodedData.ReadByte()
		if err != nil {
			return nil, err
		}

		dataBytes = append(dataBytes, b)
	}

	return string(dataBytes), nil
}

func decodeInteger(encodedData *bufio.Reader) (data interface{}, err error) {
	dataBytes := make([]byte, 0)
	for {
		b, err := encodedData.ReadByte()
		if err != nil {
			return nil, err
		} else if b == 'e' {
			break
		}

		dataBytes = append(dataBytes, b)
	}

	dataStr := string(dataBytes)
	if dataStr == "-0" {
		return nil, ErrMinusZeroInteger
	} else if len(dataStr) > 1 && strings.ReplaceAll(dataStr, "-", "")[0] == '0' {
		return nil, NewErrLeadingZeroInteger(dataStr)
	}

	return strconv.Atoi(dataStr)
}

func decodeList(encodedData *bufio.Reader) (data interface{}, err error) {
	list := make([]interface{}, 0)

	for {
		b, err := encodedData.ReadByte()
		if err != nil {
			return nil, err
		} else if b == 'e' {
			break
		}

		encodedData.UnreadByte()
		item, err := decodeNext(encodedData)
		if err != nil {
			return nil, err
		}

		list = append(list, item)
	}

	return list, nil
}

func decodeDictionary(encodedData *bufio.Reader) (data interface{}, err error) {
	dict := make(map[string]interface{})

	for {
		b, err := encodedData.ReadByte()
		if err != nil {
			return nil, err
		} else if b == 'e' {
			break
		}

		encodedData.UnreadByte()
		rawKey, err := decodeNext(encodedData)
		if err != nil {
			return nil, err
		}

		key, ok := rawKey.(string)
		if !ok {
			return nil, fmt.Errorf("a key in a dictionary must be a string, %v doest not satisfy", rawKey)
		}

		item, err := decodeNext(encodedData)
		if err != nil {
			return nil, err
		}

		dict[key] = item
	}

	return dict, nil
}
