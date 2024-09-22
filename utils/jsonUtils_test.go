package utils

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestReadBody(t *testing.T) {
	tests := []struct {
		body         string
		expected     map[string]interface{}
		expectError  bool
	}{
		{"{\"key\":\"value\"}", map[string]interface{}{"key": "value"}, false},
		{"invalid json", nil, true},
	}

	for _, test := range tests {
		req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(test.body))
		var result map[string]interface{}

		err := ReadBody(req, &result)

		if test.expectError && err.Err == nil {
			t.Errorf("Expected error for body %q, got nil", test.body)
		} else if !test.expectError && err.Err != nil {
			t.Errorf("Did not expect error for body %q, got %v", test.body, err.Err)
		} else if !test.expectError && !compareMaps(result, test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestReadBodyMultipart(t *testing.T) {
	tests := []struct {
		formValue    string
		expected     map[string]interface{}
		expectError  bool
	}{
		{"{\"key\":\"value\"}", map[string]interface{}{"key": "value"}, false},
		{"invalid json", nil, true},
	}

	for _, test := range tests {
		req := httptest.NewRequest("POST", "/test", nil)
		req.Form = make(map[string][]string)
		req.Form.Add("form_content", test.formValue)
		var result map[string]interface{}

		err := ReadBodyMultipart(req, &result)

		if test.expectError && err.Err == nil {
			t.Errorf("Expected error for form_value %q, got nil", test.formValue)
		} else if !test.expectError && err.Err != nil {
			t.Errorf("Did not expect error for form_value %q, got %v", test.formValue, err.Err)
		} else if !test.expectError && !compareMaps(result, test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func compareMaps(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valA := range a {
		valB, exists := b[key]
		if !exists || valA != valB {
			return false
		}
	}
	return true
}