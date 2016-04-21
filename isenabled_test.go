package debugflag_test

import (
	"github.com/andreyvit/debugflag"
	"strings"
	"testing"
)

var enabledTests = []struct {
	spec   string
	inputs string
}{
	{"", "-foo -foo:bar -bar"},
	{"foo", "+foo +foo:bar -bar"},
	{"foo:bar", "-foo +foo:bar -bar"},
	{"-foo,foo:bar", "-foo +foo:bar -bar"},
	{"foo,-foo:bar", "+foo -foo:bar -foo:bar:boz +foo:boz -bar"},
	{"all", "+foo +foo:bar +bar"},
}

func TestIsEnabled(t *testing.T) {
	for _, tt := range enabledTests {
		debugflag.Reset(tt.spec)

		for _, input := range strings.Split(tt.inputs, " ") {
			expected := strings.HasPrefix(input, "+")
			if !strings.HasPrefix(input, "+") && !strings.HasPrefix(input, "-") {
				t.Fatalf("Invalid input %q in test for spec %q", input, tt.spec)
			}
			flag := input[1:]

			actual := debugflag.IsEnabled(flag)
			if actual != expected {
				t.Errorf("<%s>.IsEnabled(%q) == %q, expected %q", tt.spec, flag, actual, expected)
			} else {
				t.Logf("<%s>.IsEnabled(%q) == %q", tt.spec, flag, actual)
			}
		}
	}
}
