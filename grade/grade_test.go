package grade

import (
	"fmt"
	"testing"
)

func ptr(f float32) *float32 {
	return &f
}

func TestGradeInsert(t *testing.T) {
	tests := []struct {
		grades      []float32
		expectError bool
	}{
		{grades: []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expectError: false},
		{grades: []float32{-5, 12}, expectError: true},
		{grades: []float32{}, expectError: true},
	}

	for _, tt := range tests {
		var result []float32
		err := InsertGrades(&result, tt.grades)

		fmt.Printf("Result: %v\nExpected: %v", result, tt.grades)

		if tt.expectError && err == nil {
			t.Errorf("expected error %v, but wasnt thrown", tt.grades)
		}
		if !tt.expectError && err != nil {
			t.Errorf("did not expect error %v, but caught: %v", tt.grades, err)
		}
	}
}

func TestCalculateGradeAvg(t *testing.T) {
	tests := []struct {
		name     string
		input    []float32
		expected *float32
	}{
		{"three grades", []float32{10, 8, 6}, ptr(8)},
		{"two grades", []float32{7, 9}, ptr(8)},
		{"one grade", []float32{10}, ptr(10)},
		{"no grade", []float32{}, nil},
		{"negative grades", []float32{-5, -12}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateGradeAvg(tt.input)

			if tt.expected == nil {
				fmt.Printf("Grades: %v, Result: nil\n", tt.input)
			} else {
				fmt.Printf("Grades: %v, Result: %.2f\n", tt.input, *tt.expected)
			}

			if tt.expected == nil && result != nil {
				t.Errorf("expected nil, got %.2f", *result)
				return
			}

			if tt.expected != nil && result == nil {
				t.Errorf("expected %.2f, got nil", *tt.expected)
				return
			}

			if tt.expected != nil && result != nil && *result != *tt.expected {
				t.Errorf("expected %.2f, got %.2f", *tt.expected, *result)
			}
		})
	}
}
