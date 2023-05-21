package peersmsg

import (
	"fmt"
	"strings"
	"testing"
)

// TestFromString checks if Parser correctly parses valid strings into messages
// and returns appropriate errors for invalid strings.
func TestFromString(t *testing.T) {
	// Set up a new Parser with a simple validator that rejects any message containing "invalid"
	p := NewParser('\n', WithValidator(func(m Message) error {
		if strings.Contains(m.String(), "invalid") {
			return fmt.Errorf("validation error: invalid message")
		}
		return nil
	}))

	// Define some test cases
	testCases := []struct {
		input          string
		shouldError    bool
		expectedOutput string
	}{
		{
			input:          "valid message",
			shouldError:    false,
			expectedOutput: "valid message",
		},
		{
			input:       "invalid message",
			shouldError: true,
		},
	}

	// Run each test case
	for _, tt := range testCases {
		result, err := p.(*ParserRaw).FromString(tt.input)
		// If the test case should error, check that an error was returned
		if tt.shouldError {
			if err == nil {
				t.Errorf("expected error, got nil")
			}
		} else {
			// Otherwise, check that no error was returned and the output is as expected
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if result.String() != tt.expectedOutput {
				t.Errorf("expected output %q, got %q", tt.expectedOutput, result.String())
			}
		}
	}
}

// TestNextMessage checks that Parser correctly reads messages from a reader and applies the validator.
func TestNextMessage(t *testing.T) {
	// Set up a reader with a valid message and an invalid message
	reader := strings.NewReader("valid message\ninvalid message")
	p := NewParser('\n', WithValidator(func(m Message) error {
		if strings.Contains(m.String(), "invalid") {
			return fmt.Errorf("validation error: invalid message")
		}
		return nil
	}))

	// Test that the first message is read correctly
	result, err := p.NextMessage(reader)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.String() != "valid message" {
		t.Errorf("expected output 'valid message', got %q", result.String())
	}

	// Test that the second, invalid message causes an error
	result, err = p.NextMessage(reader)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

// TestSetValidator checks that the validator is correctly set in the parser
func TestSetValidator(t *testing.T) {
	p := &ParserRaw{}

	validator := func(m Message) error { return nil }
	p.SetValidator(validator)
	if p.validator == nil {
		t.Errorf("validator should not be nil after being set")
	}
}

// TestSetLogger checks that the logger is correctly set in the parser
func TestSetLogger(t *testing.T) {
	p := &ParserRaw{}

	logger := func(args ...interface{}) { fmt.Println(args...) }
	p.SetLogger(logger)
	if p.logger == nil {
		t.Errorf("logger should not be nil after being set")
	}
}
