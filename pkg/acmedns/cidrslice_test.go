package acmedns

import (
	"encoding/json"
	"testing"
)

func TestCidrSlice(t *testing.T) {
	for i, test := range []struct {
		input       Cidrslice
		expectedErr bool
		expectedLen int
	}{
		{[]string{"192.168.1.0/24"}, false, 1},
		{[]string{"shoulderror"}, true, 0},
		{[]string{"2001:db8:aaaaa::"}, true, 0},
		{[]string{"192.168.1.0/24", "2001:db8::/32"}, false, 2},
	} {
		err := test.input.IsValid()
		if test.expectedErr && err == nil {
			t.Errorf("Expected test %d to generate IsValid() error but it didn't", i)
		}
		if !test.expectedErr && err != nil {
			t.Errorf("Expected test %d to pass IsValid() but it generated an error %s", i, err)
		}
		outSlice := []string{}
		err = json.Unmarshal([]byte(test.input.JSON()), &outSlice)
		if err != nil {
			t.Errorf("Unexpected error when unmarshaling Cidrslice JSON: %s", err)
		}
		if len(outSlice) != test.expectedLen {
			t.Errorf("Expected cidrslice JSON to be of length %d, but got %d instead for test %d", test.expectedLen, len(outSlice), i)
		}
	}
}

func TestAllowedFrom(t *testing.T) {
	testslice := Cidrslice([]string{"192.168.1.0/24", "2001:db8::/32"})
	for _, test := range []struct {
		input    string
		expected bool
	}{
		{"192.168.1.42", true},
		{"192.168.2.42", false},
		{"2001:db8:aaaa::", true},
		{"2001:db9:aaaa::", false},
	} {
		if testslice.AllowedFrom(test.input) != test.expected {
			t.Errorf("Was expecting AllowedFrom to return %t for %s but got %t instead.", test.expected, test.input, !test.expected)
		}
	}
}

func TestAllowedFromList(t *testing.T) {
	testslice := Cidrslice([]string{"192.168.1.0/24", "2001:db8::/32"})
	if testslice.AllowedFromList([]string{"192.168.2.2", "1.1.1.1"}) != false {
		t.Errorf("Was expecting AllowedFromList to return false")
	}
	if testslice.AllowedFromList([]string{"192.168.1.2", "1.1.1.1"}) != true {
		t.Errorf("Was expecting AllowedFromList to return true")
	}
	allowfromall := Cidrslice([]string{})
	if allowfromall.AllowedFromList([]string{"192.168.1.2", "1.1.1.1"}) != true {
		t.Errorf("Expected non-restricted AllowFrom to be allowed")
	}
	if allowfromall.AllowedFromList([]string{}) != true {
		t.Errorf("Expected non-restricted AllowFrom to be allowed for empty list")
	}
}
