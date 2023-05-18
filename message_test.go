package peersmsg

import (
	"errors"
	"strings"
	"testing"
)

func TestNextMessage(t *testing.T) {
	data := "hello\nworld\n"
	reader := strings.NewReader(data)
	parser := NewParser('\n')

	message, err := parser.NextMessage(reader)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if message.String() != "hello" {
		t.Fatalf("expected %q, got %q", "hello", message.String())
	}

	// Expecting EOF error here
	_, err = parser.NextMessage(reader)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestNextMessage_CustomSeparator(t *testing.T) {
	data := "hello|world|"
	reader := strings.NewReader(data)
	parser := NewParser('|')

	message, err := parser.NextMessage(reader)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if message.String() != "hello" {
		t.Fatalf("expected %q, got %q", "hello", message.String())
	}

	// Expecting EOF error here
	_, err = parser.NextMessage(reader)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestNextMessage_WithoutTrailingSeparator(t *testing.T) {
	data := "hello\nworld"
	reader := strings.NewReader(data)
	parser := NewParser('\n')

	message, err := parser.NextMessage(reader)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if message.String() != "hello" {
		t.Fatalf("expected %q, got %q", "hello", message.String())
	}

	// Expecting EOF error here
	_, err = parser.NextMessage(reader)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestNextMessage_WithValidator(t *testing.T) {
	data := "hello\nworld\n"
	reader := strings.NewReader(data)
	validator := func(m Message) error {
		if m.String() == "hello" {
			return errors.New("invalid message")
		}
		return nil
	}
	parser := NewParser('\n', WithValidator(validator), WithLogger(func(i ...interface{}) {
		// no logger
	}))

	_, err := parser.NextMessage(reader) // world
	if err == nil || err.Error() != "invalid message" {
		t.Fatalf("expected validation error, got %v", err)
	}
}
