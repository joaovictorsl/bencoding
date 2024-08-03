package bencoding_test

import (
	"errors"
	"testing"

	"github.com/joaovictorsl/bencoding"
)

func TestEncodeInvalidType(t *testing.T) {
	// Setup
	type invalid struct {
		Name string
	}
	invalidType := invalid{Name: "invalidType"}
	// Action
	encodedString, err := bencoding.Encode(invalidType)
	// Assert
	expectedErr := bencoding.NewErrInvalidType(invalidType)
	if err == nil {
		t.Error("expected error, got nil")
	} else if !errors.Is(expectedErr, err) {
		t.Errorf("expected error to be \"%s\", got \"%s\"", expectedErr, err)
	} else if encodedString != "" {
		t.Errorf("expected encodedString to be empty when an error happens, got %s", encodedString)
	}
}

func TestEncodeDictionary(t *testing.T) {
	// Setup
	okCases := []struct {
		input  map[string]interface{}
		output string
	}{
		{
			map[string]interface{}{
				"1": map[string]interface{}{
					"test": "nice",
				},
				"first": []interface{}{"a", "b", "c"},
				"oi":    "tchau",
				"sqrt4": 2,
			},
			"d1:1d4:test4:nicee5:firstl1:a1:b1:ce2:oi5:tchau5:sqrt4i2ee",
		},
		{
			map[string]interface{}{
				"2": map[string]interface{}{
					"good": "morning",
				},
				"hi":     "bye",
				"second": []interface{}{"d", "e", "f"},
				"sqrt16": 4,
			},
			"d1:2d4:good7:morninge2:hi3:bye6:secondl1:d1:e1:fe6:sqrt16i4ee",
		},
		{
			map[string]interface{}{
				"3": map[string]interface{}{
					"come": "here",
				},
				"different": "equal",
				"sqrt64":    8,
				"third":     []interface{}{"g", "h", "i"},
			},
			"d1:3d4:come4:heree9:different5:equal6:sqrt64i8e5:thirdl1:g1:h1:iee",
		},
	}
	for _, okCase := range okCases {
		// Action
		encodedData, err := bencoding.Encode(okCase.input)
		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		} else if encodedData != okCase.output {
			t.Errorf("expected encodedData to be %s, got %s", okCase.output, encodedData)
		}
	}
}

func TestEncodeList(t *testing.T) {
	// Setup
	okCases := []struct {
		input  []interface{}
		output string
	}{
		{
			[]interface{}{
				"testing",
				2,
				map[string]interface{}{"a": 1},
				[]interface{}{"pineapple", 0},
			},
			"l7:testingi2ed1:ai1eel9:pineapplei0eee",
		},
		{
			[]interface{}{
				"pineapple",
				0,
				map[string]interface{}{"b": 2},
				[]interface{}{"escape", -2},
			},
			"l9:pineapplei0ed1:bi2eel6:escapei-2eee",
		},
		{
			[]interface{}{
				"escape",
				-2,
				map[string]interface{}{"c": 3},
				[]interface{}{"testing", 2},
			},
			"l6:escapei-2ed1:ci3eel7:testingi2eee",
		},
	}
	for _, okCase := range okCases {
		// Action
		encodedData, err := bencoding.Encode(okCase.input)
		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		} else if encodedData != okCase.output {
			t.Errorf("expected encodedData to be %s, got %s", okCase.output, encodedData)
		}
	}
}

func TestEncodeInteger(t *testing.T) {
	// Setup
	okCases := []struct {
		input  int
		output string
	}{
		{-2, "i-2e"},
		{2, "i2e"},
		{0, "i0e"},
	}
	for _, okCase := range okCases {
		// Action
		encodedData, err := bencoding.Encode(okCase.input)
		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		} else if encodedData != okCase.output {
			t.Errorf("expected encodedData to be %s, got %s", okCase.output, encodedData)
		}
	}
}

func TestEncodeString(t *testing.T) {
	// Setup
	okCases := []struct {
		input  string
		output string
	}{
		{"example", "7:example"},
		{"a", "1:a"},
		{"bit", "3:bit"},
		{"olympics", "8:olympics"},
		{"", ""},
	}
	for _, okCase := range okCases {
		// Action
		encodedData, err := bencoding.Encode(okCase.input)
		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		} else if encodedData != okCase.output {
			t.Errorf("expected encodedData to be %s, got %s", okCase.output, encodedData)
		}
	}
}
