package fib

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	for _, tt := range testCases {
		res := Fibonacci(tt.input)

		if res != tt.expected {
			t.Errorf("FAIL: Expected %d but got %d", tt.expected, res)
			return
		}

		t.Logf("PASS: Expected %d and got %d", tt.expected, res)
	}
}

func TestMemoizedFibonacci(t *testing.T) {
	for _, tt := range testCases {
		fn := newMemoizedFib()
		res := fn(tt.input)

		if res != tt.expected {
			t.Errorf("FAIL: Expected %d but got %d", tt.expected, res)
			return
		}

		t.Logf("PASS: Expected %d and got %d", tt.expected, res)
	}
}
