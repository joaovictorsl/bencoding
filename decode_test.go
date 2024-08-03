package bencoding_test

import (
	"bufio"
	"bytes"
	"errors"
	"maps"
	"reflect"
	"slices"
	"testing"

	"github.com/joaovictorsl/bencoding"
)

func TestDecodeDictionary(t *testing.T) {
	// Setup
	okCases := []struct {
		v       map[string]interface{}
		encoded []byte
	}{
		{map[string]interface{}{"test": "torrent"}, []byte("d4:test7:torrente")},
		{map[string]interface{}{"torrent": "peer"}, []byte("d7:torrent4:peere")},
		{map[string]interface{}{"peer": "bit"}, []byte("d4:peer3:bite")},
		{map[string]interface{}{"bit": "past"}, []byte("d3:bit4:paste")},
		{map[string]interface{}{"bit": 32}, []byte("d3:biti32ee")},
	}

	for _, okCase := range okCases {
		// Action
		data, err := bencoding.Decode(bufio.NewReader(bytes.NewReader(okCase.encoded)))
		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		dataMap, ok := data.(map[string]interface{})
		if !ok {
			t.Errorf("expected data to be a map[string]interface {}, got %v", reflect.TypeOf(data))
		} else if !maps.Equal(dataMap, okCase.v) {
			t.Errorf("expected data to be %v, got %v", okCase.v, dataMap)
		}
	}
}

func TestDecodeList(t *testing.T) {
	// Setup
	okCases := []struct {
		v       []interface{}
		encoded []byte
	}{
		{[]interface{}{"test", "torrent"}, []byte("l4:test7:torrente")},
		{[]interface{}{"torrent", "peer"}, []byte("l7:torrent4:peere")},
		{[]interface{}{"peer", "bit"}, []byte("l4:peer3:bite")},
		{[]interface{}{"bit", "past"}, []byte("l3:bit4:paste")},
	}

	for _, okCase := range okCases {
		// Action
		data, err := bencoding.Decode(bufio.NewReader(bytes.NewReader(okCase.encoded)))
		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		dataSlice, ok := data.([]interface{})
		if !ok {
			t.Errorf("expected data to be a []interface {}, got %v", reflect.TypeOf(data))
		} else if !slices.Equal(dataSlice, okCase.v) {
			t.Errorf("expected data to be %v, got %v", okCase.v, dataSlice)
		}
	}
}

func TestDecodeString(t *testing.T) {
	// Setup
	okCases := []struct {
		v       string
		encoded []byte
	}{
		{"test", []byte("4:test")},
		{"torrent", []byte("7:torrent")},
		{"peer", []byte("4:peer")},
		{"bit", []byte("3:bit")},
	}

	for _, okCase := range okCases {
		// Action
		data, err := bencoding.Decode(bufio.NewReader(bytes.NewReader(okCase.encoded)))
		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		dataStr, ok := data.(string)
		if !ok {
			t.Errorf("expected data to be a string, got %v", reflect.TypeOf(data))
		} else if dataStr != okCase.v {
			t.Errorf("expected data to be %v, got %v", okCase.v, dataStr)
		}
	}
}

func TestDecodeInteger(t *testing.T) {
	// Setup
	okCases := []struct {
		v       int
		encoded []byte
	}{
		{23, []byte("i23e")},
		{32, []byte("i+32e")},
		{-23, []byte("i-23e")},
		{0, []byte("i0e")},
	}

	for _, okCase := range okCases {
		// Action
		data, err := bencoding.Decode(bufio.NewReader(bytes.NewReader(okCase.encoded)))
		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		dataInteger, ok := data.(int)
		if !ok {
			t.Errorf("expected data to be an integer, got %v", reflect.TypeOf(data))
		} else if dataInteger != okCase.v {
			t.Errorf("expected data to be %d, got %d", okCase.v, dataInteger)
		}
	}
}

func TestDecodeIntegerIncomplete(t *testing.T) {
	// Setup
	errCases := []struct {
		v       int
		encoded []byte
	}{
		{23, []byte("i23")},
		{32, []byte("i+32")},
		{-23, []byte("i-23")},
		{0, []byte("i0")},
	}

	for _, errCase := range errCases {
		// Action
		data, err := bencoding.Decode(bufio.NewReader(bytes.NewReader(errCase.encoded)))
		// Assert
		if err == nil {
			t.Error("expected error, got nil")
		} else if data != nil {
			t.Errorf("expected data to be nil when error happens, got %v", data)
		}
	}
}

func TestDecodeIntegerLeadingZero(t *testing.T) {
	// Setup
	errCases := []struct {
		v       int
		encoded []byte
	}{
		{3, []byte("i03e")},
		{-3, []byte("i-03e")},
	}

	for _, errCase := range errCases {
		data, err := bencoding.Decode(bufio.NewReader(bytes.NewReader(errCase.encoded)))
		if err == nil {
			t.Error("expected error, got nil")
		} else if !errors.Is(bencoding.ErrLeadingZeroInteger, err) {
			t.Errorf("expected error to be %v, got %v", bencoding.ErrLeadingZeroInteger, err)
		} else if data != nil {
			t.Errorf("expected data to be nil when error happens, got %v", data)
		}
	}
}

func TestDecodeIntegerMinusZero(t *testing.T) {
	data, err := bencoding.Decode(bufio.NewReader(bytes.NewReader([]byte("i-0e"))))
	if err == nil {
		t.Error("expected error, got nil")
	} else if !errors.Is(bencoding.ErrMinusZeroInteger, err) {
		t.Errorf("expected error to be %v, got %v", bencoding.ErrMinusZeroInteger, err)
	} else if data != nil {
		t.Errorf("expected data to be nil when error happens, got %v", data)
	}
}
