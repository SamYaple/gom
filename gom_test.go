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

func TestTuple(t *testing.T) {
	parser := Tuple(Tag("Hello"), Tag(" "), Tag("World"))
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

func TestAlphanumeric1(t *testing.T) {
	parser := Alphanumeric1()
	input := "abc123!xyz"

	remaining, result, err := parser(input)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != "abc123" {
		t.Errorf("Expected result 'abc123', got: %v", result)
	}
	if remaining != "!xyz" {
		t.Errorf("Expected remaining '!xyz', got: %s", remaining)
	}
}

func TestAlphanumeric1_NoMatch(t *testing.T) {
	parser := Alphanumeric1()
	input := "!xyz"

	_, _, err := parser(input)
	if err == nil {
		t.Error("Expected an error because input does not start with an alphanumeric")
	}
}

func TestMany1(t *testing.T) {
	parser := Many1(Alphanumeric1())
	input := "abc123 xyz789"

	remaining, result, err := parser(input)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	results, ok := result.([]interface{})
	if !ok {
		t.Fatalf("Expected []interface{}, got %T", result)
	}

	if len(results) != 1 {
		t.Errorf("Expected 1 match, got %d", len(results))
	}
	if results[0] != "abc123" {
		t.Errorf("Expected 'abc123', got %v", results[0])
	}
	if remaining != " xyz789" {
		t.Errorf("Expected remaining ' xyz789', got %s", remaining)
	}
}

func TestMany1_NoMatch(t *testing.T) {
	parser := Many1(Alphanumeric1())
	input := "# no alpha at start"

	_, _, err := parser(input)
	if err == nil {
		t.Error("Expected an error because the first parser call fails immediately")
	}
}

func TestRecognize(t *testing.T) {
	parser := Recognize(Alphanumeric1())
	input := "abc123!xyz"

	remaining, result, err := parser(input)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != "abc123" {
		t.Errorf("Expected recognized substring 'abc123', got %v", result)
	}
	if remaining != "!xyz" {
		t.Errorf("Expected remaining '!xyz', got %s", remaining)
	}
}
