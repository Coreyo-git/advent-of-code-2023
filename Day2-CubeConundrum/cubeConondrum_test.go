package main

import "testing"

type addTest struct {
    arg1, arg2, arg3, expected int
}

var addTests = []addTest{
    addTest{4, 2, 6, 48},
    addTest{1, 3, 4, 12},
    addTest{20, 13, 6, 1560},
    addTest{14, 3, 15, 630},
    addTest{6, 3, 2, 36},
}

func TestPowers(t *testing.T) {
	var expectedTotal = 2286
	var outputTotal = 0

	for _, test := range addTests {
        if output := test.arg1 * test.arg2 * test.arg3; output != test.expected {
            t.Errorf("Output %q not equal to expected %q", output, test.expected)
        }
		outputTotal += test.arg1 * test.arg2 * test.arg3
    }
	if outputTotal != expectedTotal {
		t.Errorf("Output Total %q not equal to expected Total %q", outputTotal, expectedTotal)
	}

}

func TestAdd(t *testing.T){

    got := Add(4, 6)
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetStringAsInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Game 1", 1},
		{"Game 42", 42},
		{"7 red", 7},
		{"No numeric in string", 0},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := getStringAsInt(test.input)

			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestIsDrawWithinBounds(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"5 red", true},    // within bounds
		{"15 red", false},  // exceeds bounds
		{"8 blue", true},   // within bounds
		{"18 blue", false}, // exceeds bounds
		{"10 green", true}, // within bounds
		{"20 green", false},// exceeds bounds
		{"invalid", false}, // invalid color
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := isDrawWithinBounds(test.input)

			if result != test.expected {
				t.Errorf("Expected %t, got %t", test.expected, result)
			}
		})
	}
}