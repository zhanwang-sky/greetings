package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Jon Snow"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Errorf("Hello(%s) returns %q, %v, want match for %#q, nil", name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") returns %q, %v, want "", error`, msg, err)
	}
}
