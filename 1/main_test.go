package main

import "testing"

func Test_part1(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expected := 142
	actual := part1(input)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func Test_part2(t *testing.T) {
	t.Error("No tests written")
}
