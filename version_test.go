package version

import (
	"fmt"
	"runtime"
	"testing"
)

func TestVersionString(t *testing.T) {
	tests := []struct {
		label    string
		input    Info
		actual   string
		expected string
	}{
		{
			label:    "package default",
			input:    Version,
			expected: fmt.Sprintf("unknown v0-dev (unknown) %v/%v - BuildDate: unknown", runtime.GOOS, runtime.GOARCH),
			actual:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			test.actual = test.input.String()

			if test.expected != test.actual {
				t.Fatalf("expected: %v\n\ngot: %v\n\n", test.expected, test.actual)
			}
		})
	}
}
