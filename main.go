// // Successful test
// package main

// // Add returns a + b
// func Add(a, b int) int {
// 	return a + b
// }

// unsuccessfull tests
package main

import "testing"

func TestSample(t *testing.T) {
	// Force a failure
	got := 1
	want := 2

	if got != want {
		t.Fatalf("expected %d but got %d ❌", want, got)
	}
}
