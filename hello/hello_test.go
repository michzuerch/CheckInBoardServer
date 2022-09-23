package hello

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Sample for unittest" {
		t.Errorf("Greet() = %s; want Sample for unittest", result)
	}
}
