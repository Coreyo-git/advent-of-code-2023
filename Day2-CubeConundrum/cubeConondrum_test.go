package main

import "testing"

type addTest struct {
	arg1 string 
	expected int
}

var addTests = []addTest{
	addTest{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1},
	addTest{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 2},
	addTest{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 0},
	addTest{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 0},
	addTest{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 5},
}

func Test(t *testing.T){

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