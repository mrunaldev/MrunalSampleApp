package main

import (
    "testing"
)

func TestCalculator_CalculateSquare(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"Zero", 0, 0},
        {"Positive", 2, 4},
        {"Negative", -3, 9},
    }

    calc := &Calculator{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := calc.CalculateSquare(tt.input); got != tt.expected {
                t.Errorf("CalculateSquare() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestCalculator_CalculateCube(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"Zero", 0, 0},
        {"Positive", 2, 8},
        {"Negative", -3, -27},
    }

    calc := &Calculator{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := calc.CalculateCube(tt.input); got != tt.expected {
                t.Errorf("CalculateCube() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestCalculator_CalculateConcurrently(t *testing.T) {
    tests := []struct {
        name          string
        input         int
        expectedSquare int
        expectedCube   int
    }{
        {"Zero", 0, 0, 0},
        {"Positive", 2, 4, 8},
        {"Negative", -3, 9, -27},
    }

    calc := &Calculator{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            square, cube := calc.CalculateConcurrently(tt.input)
            if square != tt.expectedSquare {
                t.Errorf("CalculateConcurrently() square = %v, want %v", square, tt.expectedSquare)
            }
            if cube != tt.expectedCube {
                t.Errorf("CalculateConcurrently() cube = %v, want %v", cube, tt.expectedCube)
            }
        })
    }
}