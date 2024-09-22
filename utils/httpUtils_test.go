package utils

import (
	"net/http/httptest"
	"testing"
)

func TestReadIDParam(t *testing.T) {
	tests := []struct {
		idParam    string
		expectedID int
	}{
		{"1", 1},
		{"42", 42},
		{"100", 100},
		{"0", 0},
		{"-1", -1},
		{"notanumber", 0}, // Tidak menangani error
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/item?ID="+test.idParam, nil)

		result := ReadIDParam(req)

		if result != test.expectedID {
			t.Errorf("ReadIDParam(%q) = %d; want %d", test.idParam, result, test.expectedID)
		}
	}
}