package plain

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	x := Sum(2, 3)
	if x != 5 {
		t.Error("Expected ", 5, " Got ", x)
	}
}

func ExampleSum() {

	fmt.Println(Sum(2,3))
	// Output:
	// 5
}
