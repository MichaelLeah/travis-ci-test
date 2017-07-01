package travis

import "testing"

func TestDouble(t *testing.T) {
	num := Double(5)
	if num != 10 {
		t.Error("Expected 10, got ", num)
	}
}
