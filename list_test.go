package main

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Arrange
	expected := 0
	// Act
	result := NewList()
	// Assert
	if result.len != expected {
		t.Errorf("Incorrect result. Expect %d, got %d", expected, result.len)
	}
}
