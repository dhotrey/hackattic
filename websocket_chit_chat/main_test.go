package main

import "testing"

func TestGetClosestDuration(t *testing.T) {
	closest := GetClosestDuration(int64(1450))
	expected := int64(1500)

	if closest != expected {
		t.Errorf("got %d , expected %d", closest, expected)
	}
}
