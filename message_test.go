package peersmsg

import (
	"reflect"
	"testing"
)

func TestMessageRaw_Bytes(t *testing.T) {
	data := []byte("hello")
	m := MessageRaw{Data: data}

	result := m.Bytes()

	if !reflect.DeepEqual(result, data) {
		t.Errorf("expected %v, got %v", data, result)
	}
}

func TestMessageRaw_String(t *testing.T) {
	data := []byte("hello")
	m := MessageRaw{Data: data}

	expected := "hello"
	result := m.String()

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
