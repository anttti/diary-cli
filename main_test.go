package main

import (
	"regexp"
	"testing"
)

func TestFormatEntry(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		pattern string
	}{
		{
			name:    "simple text",
			text:    "This is a diary entry",
			pattern: `^\[\d{2}:\d{2}:\d{2}\] This is a diary entry\n$`,
		},
		{
			name:    "text with special characters",
			text:    "Entry with !@#$%^&*() special chars",
			pattern: `^\[\d{2}:\d{2}:\d{2}\] Entry with !@#\$%\^&\*\(\) special chars\n$`,
		},
		{
			name:    "empty text",
			text:    "",
			pattern: `^\[\d{2}:\d{2}:\d{2}\] \n$`,
		},
		{
			name:    "multiline text",
			text:    "Line 1\nLine 2",
			pattern: `^\[\d{2}:\d{2}:\d{2}\] Line 1\nLine 2\n$`,
		},
		{
			name:    "text with tabs",
			text:    "Text	with	tabs",
			pattern: `^\[\d{2}:\d{2}:\d{2}\] Text	with	tabs\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatEntry(tt.text)

			matched, err := regexp.MatchString(tt.pattern, result)
			if err != nil {
				t.Fatalf("regex error: %v", err)
			}

			if !matched {
				t.Errorf("formatEntry(%q) = %q, doesn't match pattern %q", tt.text, result, tt.pattern)
			}
		})
	}
}
