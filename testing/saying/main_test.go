package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {

	s := Greet("Karan")

	if s != "Welcome my dear Karan" {
		t.Error("Got", s, "Expected", "Welcome my dear Karan")
	}
}

func ExampleGreet() {

	fmt.Println(Greet("Karan"))
	// Output:
	// Welcome my dear Karan
}

func BenchmarkGreet(b *testing.B) { //go test -bench ./Greet

	for i := 0; i < b.N; i++ {
		Greet("Karan")
	}
}
