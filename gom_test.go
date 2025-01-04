package gom

import (
	"testing"
)

func TestTag(t *testing.T) {
	parser := Tag("Hello")
	input := "Hello World"

	remaining, result, err := parser(input)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != "Hello" {
		t.Errorf("Expected result 'Hello', got: %v", result)
	}
	if remaining != " World" {
		t.Errorf("Expected remaining ' World', got: %s", remaining)
	}
}

func TestTagNoMatch(t *testing.T) {
	parser := Tag("Hello")
	input := "Hi World"

	remaining, result, err := parser(input)
	if err == nil {
		t.Fatal("Expected an error, got none")
	}
	if remaining != "Hi World" {
		t.Errorf("Expected input to remain 'Hi World' on failure, got: %s", remaining)
	}
	if result != nil {
		t.Errorf("Expected result to be nil on failure, got: %v", result)
	}
}

func TestSequence(t *testing.T) {
	parser := Sequence(Tag("Hello"), Tag(" "), Tag("World"))
	input := "Hello World"

	remaining, result, err := parser(input)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	matches, ok := result.([]interface{})
	if !ok {
		t.Fatalf("Expected result to be []interface{}, got something else")
	}
	if len(matches) != 3 {
		t.Fatalf("Expected 3 matches, got: %d", len(matches))
	}
	if matches[0] != "Hello" || matches[1] != " " || matches[2] != "World" {
		t.Errorf("Unexpected matches: %v", matches)
	}
	if remaining != "" {
		t.Errorf("Expected empty remaining string, got: %s", remaining)
	}
}

func TestAlt(t *testing.T) {
	parser := Alt(Tag("Hello"), Tag("Hi"))
	input := "Hi World"

	remaining, result, err := parser(input)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != "Hi" {
		t.Errorf("Expected result 'Hi', got: %v", result)
	}
	if remaining != " World" {
		t.Errorf("Expected remaining ' World', got: %s", remaining)
	}
}
