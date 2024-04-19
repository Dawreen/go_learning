package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName call greetings.Hello with a name an checks
// if it works
func TestHelloName(t *testing.T) {
	name := "Jimmy"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Jimmy")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Jimmy") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
