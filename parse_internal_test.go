package debugflag

import (
	"strings"
	"testing"
)

var parseTests = []struct {
	input    string
	expected string
}{
	{"", ""},
	{"   ", ""},
	{"foo", "+foo"},
	{"  foo  ", "+foo"},
	{"  !foo  ", "-foo"},
	{"  -foo  ", "-foo"},
	{"  !  foo  ", "+foo"},
	{"foo:bar", "+foo:bar"},
	{"!foo:bar", "-foo:bar"},
	{"-foo:bar", "-foo:bar"},
	{"foo,bar", "+foo +bar"},
	{"foo bar", "+foo +bar"},
	{"foo,!bar", "+foo -bar"},
	{"foo !bar", "+foo -bar"},
	{"!foo,bar", "+bar -foo"},
	{"!foo bar", "+bar -foo"},
}

func TestParse(t *testing.T) {
	for _, tt := range parseTests {
		enabled, disabled := parse(tt.input)
		actualStr := parseResultToString(enabled, disabled)
		if actualStr != tt.expected {
			t.Errorf("parse(%q) == %s, expected %s", tt.input, actualStr, tt.expected)
		} else {
			t.Logf("parse(%q) == %s", tt.input, actualStr)
		}
	}
}

func parseResultToString(enabled, disabled []string) string {
	result := make([]string, 0)
	for _, el := range enabled {
		result = append(result, "+"+el)
	}
	for _, el := range disabled {
		result = append(result, "-"+el)
	}
	return strings.Join(result, " ")
}
