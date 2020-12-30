package isdayoff

import (
	"net/http"
	"testing"
)

func TestIsLeap(t *testing.T) {
	client := New(http.DefaultClient)
	leap, err := client.IsLeap(2020)
	if err != nil {
		t.Error(err)
	}
	if leap != true {
		t.Errorf("should be true, equal: %v", leap)
	}
}
