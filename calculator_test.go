package main

import (
	"testing"
)

func TestCalculator_CalculateSquare(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expected    int
		expectError bool
	}{
		{"Zero", 0, 0, false},
		{"Positive", 2, 4, false},
		{"Negative", -3, 9, false},
		{"Large Number", 1<<30 + 1, 0, true},
	}

	calc := &Calculator{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.CalculateSquare(tt.input)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if got != tt.expected {
					t.Errorf("CalculateSquare() = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}

func TestCalculator_CalculateCube(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expected    int
		expectError bool
	}{
		{"Zero", 0, 0, false},
		{"Positive", 2, 8, false},
		{"Negative", -3, -27, false},
		{"Large Number", 1<<20 + 1, 0, true},
	}

	calc := &Calculator{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.CalculateCube(tt.input)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if got != tt.expected {
					t.Errorf("CalculateCube() = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}

func TestCalculator_CalculatePower(t *testing.T) {
	tests := []struct {
		name        string
		base        int
		power       int
		expected    int
		expectError bool
	}{
		{"Zero Base", 0, 5, 0, false},
		{"One Base", 1, 100, 1, false},
		{"Negative Base", -2, 3, -8, false},
		{"Zero Power", 5, 0, 1, false},
		{"One Power", 5, 1, 5, false},
		{"Normal Case", 2, 3, 8, false},
		{"Negative Power", 2, -1, 0, true},
		{"Large Numbers", 100, 100, 0, true},
	}

	calc := &Calculator{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.CalculatePower(tt.base, tt.power)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if got != tt.expected {
					t.Errorf("CalculatePower(%d, %d) = %v, want %v", tt.base, tt.power, got, tt.expected)
				}
			}
		})
	}
}

func TestCalculator_CalculateConcurrently(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedSquare int
		expectedCube   int
		expectError    bool
	}{
		{"Zero", 0, 0, 0, false},
		{"Positive", 2, 4, 8, false},
		{"Negative", -3, 9, -27, false},
		{"Large Number", 1<<30 + 1, 0, 0, true},
	}

	calc := &Calculator{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			square, cube, err := calc.CalculateConcurrently(tt.input)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if square != tt.expectedSquare {
					t.Errorf("CalculateConcurrently() square = %v, want %v", square, tt.expectedSquare)
				}
				if cube != tt.expectedCube {
					t.Errorf("CalculateConcurrently() cube = %v, want %v", cube, tt.expectedCube)
				}
			}
		})
	}
}
