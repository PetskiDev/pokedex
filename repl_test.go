package main

import "testing"
func TestCleanInput(t *testing.T){
	
	cases := []struct{
		input		string
		expected	[]string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: " Charmander  Bulbasaur  PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, tc := range cases{
		actual := cleanInput(tc.input)

		for i := range actual{
			if tc.expected[i] != actual[i]{
				t.Errorf("Words don't match")
			}
		}

	}
}