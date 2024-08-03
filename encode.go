package bencoding

import (
	"bytes"
	"fmt"
	"slices"
)

func Encode(data interface{}) (encodedData string, err error) {
	return encodeNext(data)
}

func encodeNext(data interface{}) (encodedData string, err error) {
	if str, ok := data.(string); ok {
		return encodeString(str)
	} else if integer, ok := data.(int); ok {
		return encodeInteger(integer)
	} else if list, ok := data.([]interface{}); ok {
		return encodeList(list)
	} else if dict, ok := data.(map[string]interface{}); ok {
		return encodeDictionary(dict)
	} else {
		return "", NewErrInvalidType(data)
	}
}

func encodeString(data string) (encodedData string, err error) {
	if len(data) == 0 {
		return "", nil
	}
	return fmt.Sprintf("%d:%s", len(data), data), nil
}

func encodeInteger(data int) (encodedData string, err error) {
	return fmt.Sprintf("i%de", data), nil
}

func encodeList(data []interface{}) (encodedData string, err error) {
	encodedData = "l"

	for _, v := range data {
		encodedV, err := encodeNext(v)
		if err != nil {
			return "", err
		}

		encodedData += encodedV
	}

	return encodedData + "e", nil
}

func encodeDictionary(data map[string]interface{}) (encodedData string, err error) {
	encodedData = "d"

	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(i, j string) int {
		return bytes.Compare([]byte(i), []byte(j))
	})

	for _, k := range keys {
		v := data[k]
		encodedK, err := encodeString(k)
		if err != nil {
			return "", err
		}

		encodedV, err := encodeNext(v)
		if err != nil {
			return "", err
		}

		encodedData += encodedK + encodedV
	}

	return encodedData + "e", nil
}
