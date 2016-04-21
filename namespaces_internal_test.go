package debugflag

import (
	"testing"
)

var namespaceTests = []struct {
	input    string
	expected string
}{
	{"", ""},
	{"  ", ""},
	{"foo", ""},
	{"foo:bar", "foo"},
	{"foo:bar:boz", "foo:bar"},
	{"  foo : bar : boz  ", "  foo : bar "},
}

func TestParentNamespace(t *testing.T) {
	for _, tt := range namespaceTests {
		actual := parentNamespace(tt.input)
		if actual != tt.expected {
			t.Errorf("parentNamespace(%q) == %q, expected %q", tt.input, actual, tt.expected)
		} else {
			t.Logf("parentNamespace(%q) == %q", tt.input, actual)
		}
	}
}

var iterateNamespacesTests = []struct {
	input    string
	expected string
}{
	{"", ""},
}

func TestIterateNamespaces(t *testing.T) {
	for _, tt := range iterateNamespacesTests {
		actual := parentNamespace(tt.input)
		actualStr := parentNamespaceResultToString(actual)
		if actualStr != tt.expected {
			t.Errorf("parentNamespace(%q) == %s, expected %s", tt.input, actualStr, tt.expected)
		} else {
			t.Logf("parentNamespace(%q) == %s", tt.input, actualStr)
		}
	}
}

func parentNamespaceResultToString(actual string) string {
	return actual
}
