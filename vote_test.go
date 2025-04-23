package vote

import "testing"

func TestHasToVote(t *testing.T) {
	tests := []struct {
		age      int
		expected bool
	}{
		{age: -4, expected: false},
		{age: 17, expected: false},
		{age: 18, expected: true},
		{age: 69, expected: true},
		{age: 70, expected: false},
	}

	for _, tt := range tests {
		result := HasToVote(tt.age)
		if result != tt.expected {
			t.Errorf("HasToVote(%d) = %v; expected %v", tt.age, result, tt.expected)
		}
	}
}
