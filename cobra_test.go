package main

import (
	"testing"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestExtractMainComment(t *testing.T) {
	tests := map[string]struct {
		input    protogen.Comments
		expected string
	}{
		"single line comment": {
			input:    protogen.Comments("// This is a comment"),
			expected: "// This is a comment",
		},
		"multi-line comment": {
			input:    protogen.Comments("// First line\n// Second line"),
			expected: "// First line",
		},
		"empty comment": {
			input:    protogen.Comments(""),
			expected: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := extractMainComment(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestExtractTrailingComments(t *testing.T) {
	tests := map[string]struct {
		input    protogen.Comments
		expected string
	}{
		"single line comment": {
			input:    protogen.Comments("// This is a comment"),
			expected: "",
		},
		"multi-line comment": {
			input:    protogen.Comments("// First line\n// Second line"),
			expected: "// Second line",
		},
		"empty comment": {
			input:    protogen.Comments(""),
			expected: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := extractTrailingComments(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestCleanComments(t *testing.T) {
	tests := map[string]struct {
		input    protogen.Comments
		expected string
	}{
		"empty comment": {
			input:    "",
			expected: "",
		},
		"single line comment": {
			input:    " This is a comment ",
			expected: "This is a comment",
		},
		"multi-line comment": {
			input:    " Line 1\n Line 2\n Line 3 ",
			expected: "Line 1\n Line 2\n Line 3",
		},
		"comment with leading and trailing spaces": {
			input:    "   Comment with spaces   ",
			expected: "Comment with spaces",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := cleanComments(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestNormalizeKind(t *testing.T) {
	tests := map[string]struct {
		input    protoreflect.Kind
		expected protoreflect.Kind
	}{
		"Sint32 to Int32":   {protoreflect.Sint32Kind, protoreflect.Int32Kind},
		"Sfixed32 to Int32": {protoreflect.Sfixed32Kind, protoreflect.Int32Kind},
		"Fixed32 to Uint32": {protoreflect.Fixed32Kind, protoreflect.Uint32Kind},
		"Sint64 to Int64":   {protoreflect.Sint64Kind, protoreflect.Int64Kind},
		"Sfixed64 to Int64": {protoreflect.Sfixed64Kind, protoreflect.Int64Kind},
		"Fixed64 to Uint64": {protoreflect.Fixed64Kind, protoreflect.Uint64Kind},
		"Group to Message":  {protoreflect.GroupKind, protoreflect.MessageKind},
		"Int32 unchanged":   {protoreflect.Int32Kind, protoreflect.Int32Kind},
		"String unchanged":  {protoreflect.StringKind, protoreflect.StringKind},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := normalizeKind(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeKind(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
