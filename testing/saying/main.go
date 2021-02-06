// Package Greet
package saying

import "fmt"

// Greet greets a person with message
func Greet(s string) string {

	return fmt.Sprint("Welcome my dear ", s)
}
