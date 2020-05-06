package tdd_in_go

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		input int
		output int
	}{
		{"0", 0, 0},
		{"2", 2, 1},
		{"10", 10, 55},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := fibonacci(test.input)
			if result != test.output {
				t.Logf("test failed. Expected: %d, got %d\n", test.input, result)
				t.Fail()
			}
		})
	}
}




