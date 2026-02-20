package urlmasker

import (
	"testing"
)

func TestURLMasker(t *testing.T) {
	testCases := []struct {
		name        string
		inputString string
		result      string
	}{
		{
			name:        "test1",
			inputString: "http://facebook.com",
			result:      "http://************",
		},
		{
			name:        "test2",
			inputString: "Hello, its my page //string//: http://localhost123.com See you http://localhost123.com",
			result:      "Hello, its my page //string//: http://**************** See you http://****************",
		},
		{
			name:        "test3",
			inputString: "http://localhost123.com Hello http http://localhost123.com :// Hello! http://localhost123.com http://localhost123.com",
			result:      "http://**************** Hello http http://**************** :// Hello! http://**************** http://****************",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			realResult := DoMaskForLink(tc.inputString)
			if realResult != tc.result {
				t.Errorf("expected %s != real %s", tc.result, realResult)
			}
		})
	}

}
