package propertbased

import "testing"

// https://codingdojo.org/kata/RomanNumerals/

func TestRomanNumerals(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ConvertToRoman(arabic int) string {
	return "I"

}
