package main

import "fmt"

//Creating a function that multiplies two integers together
func multiply(x int, y int) int {
	return x*y
}

//a unit test for the function above

package math
import "testing"

func testmultiply(t *testing.T){
	got := multiply(4,10)
	want := 40
	If got != want {
		t.ErrorF("got %q, wanted %q", got, want)
	}
}