package greetings

import (
	"regexp"  // Importing the regexp package for using regular expressions.
	"testing" // Importing the testing package which is used for writing test cases in Go.
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "LuisPa" // Assign the name "LuisPa" to be tested in the Hello function.
	// Create a regular expression that matches the exact name "LuisPa" within word boundaries.
	want := regexp.MustCompile(`\b` + name + `\b`)
	// Call the Hello function with "LuisPa" and store the message and error returned.
	msg, err := Hello("LuisPa")
	// Check if the message does not contain the exact name or if an error was returned.
	if !want.MatchString(msg) || err != nil {
		// Use t.Fatalf to report an error and immediately fail the test if the conditions are not met.
		// This method formats the string and arguments similar to fmt.Printf and then ends the test.
		// `%q` is a placeholder that quotes the string variable `msg`.
		// `%v` is a placeholder for the variable `err` that prints in the default format.
		// `%#q` is a placeholder that quotes the regexp `want` and provides its pattern in a Go-syntax representation.
		// The entire formatted message provides detailed feedback on what the test expected versus what it received,
		// making it easier to understand why the test failed.
		t.Fatalf(`Hello("LuisPa") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	// Call the Hello function with an empty string and store the results.
	msg, err := Hello("")
	// Check if a non-empty message was returned or no error was returned.
	if msg != "" || err == nil {
		// If the function does not behave as expected, fail the test and report the error.
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
