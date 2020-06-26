package main

import "testing"

func TestSqrt(t *testing.T) {
	result := sqrt()
	if result != "Code.education Rocks" {
		t.Errorf("sqrt failed: expected Code.education Rocks | got %s", result)
	}
}