package main

import (
	"testing"
)

// Test isPrime function
func TestIsPrime(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{10, false},
		{17, true},
		{19, true},
		{20, false},
	}

	for _, test := range tests {
		if result := isPrime(test.input); result != test.expected {
			t.Errorf("isPrime(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// Test isPerfect function
func TestIsPerfect(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{6, true},   // 6 = 1 + 2 + 3
		{28, true},  // 28 = 1 + 2 + 4 + 7 + 14
		{12, false}, // Not a perfect number
		{496, true}, // 496 = 1 + 2 + 4 + 8 + 16 + 31 + 62 + 124 + 248
		{8128, true},
	}

	for _, test := range tests {
		if result := isPerfect(test.input); result != test.expected {
			t.Errorf("isPerfect(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// Test isArmstrong function
func TestIsArmstrong(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{153, true},  // 1^3 + 5^3 + 3^3 = 153
		{9474, true}, // 9^4 + 4^4 + 7^4 + 4^4 = 9474
		{371, true},  // 3^3 + 7^3 + 1^3 = 371
		{407, true},  // 4^3 + 0^3 + 7^3 = 407
		{9475, false},
		{100, false},
	}

	for _, test := range tests {
		if result := isArmstrong(test.input); result != test.expected {
			t.Errorf("isArmstrong(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// Test sumOfDigits function
func TestSumOfDigits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{123, 6},  // 1 + 2 + 3 = 6
		{456, 15}, // 4 + 5 + 6 = 15
		{987, 24}, // 9 + 8 + 7 = 24
		{0, 0},    // 0 = 0
		{999, 27}, // 9 + 9 + 9 = 27
		{1111, 4}, // 1 + 1 + 1 + 1 = 4
	}

	for _, test := range tests {
		if result := sumOfDigits(test.input); result != test.expected {
			t.Errorf("sumOfDigits(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}
