package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []http.Header{
		{"Authorization": []string{"ApiKey abc123"}},
		{"Content-Type": []string{"application/json"}},
		{"Authorization": []string{"abc"}},
		{"Authorization": []string{"Bearer abc"}},
	}

	type expectedShape struct {
		S string
		E error
	}
	malformed := errors.New("malformed authorization header")
	expected := []expectedShape{
		{"abc123", nil},
		{"1", ErrNoAuthHeaderIncluded},
		{"1", malformed},
		{"1", malformed},
	}

	for i, c := range cases {
		result, err := GetAPIKey(c)
		if err != nil && !errors.Is(expected[i].E, err) {
			if expected[i].E != nil && expected[i].E.Error() != err.Error() {
				t.Fail()
			}
		}
		if expected[i].S != result {
			t.Fail()
		}
	}
}
