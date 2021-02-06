package saying

import "testing"

func TestGreet(t *testing.T) {

	s := Greet("Karan")

	if s != "Welcome my dear Karan" {
		t.Error("Got", s, "Expected", "Welcome my dear Karan")
	}
}
