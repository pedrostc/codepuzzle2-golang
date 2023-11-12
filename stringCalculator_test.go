package stringCalculator

import (
	"fmt"
	"strings"
	"testing"
)

type challengeStep struct {
	title     string
	testCases []testCase
}

type testCase struct {
	input    string
	expected int
}

func Test_challenge_steps(t *testing.T) {
	steps := []challengeStep{
		{
			title: "Empty string returns 0",
			testCases: []testCase{
				{"", 0},
			},
		},
		{
			title: "Single value returns the parsed value",
			testCases: []testCase{
				{"1", 1},
				{"2", 2},
				{"3", 3},
				{"42", 42},
				{"100", 100},
			},
		},
		{
			title: "Comma separated string returns the sum of the values",
			testCases: []testCase{
				{"1,2,3,4", 10},
				{"40,0,2", 42},
				{"1,14,7,8,12", 42},
				{"1,3,4,2,5,7,18,100", 140},
			},
		},
		{
			title: "New line separated string returns the sum of the values",
			testCases: []testCase{
				{"1\n2\n3\n4", 10},
				{"40\n0\n2", 42},
				{"1\n14\n7\n8\n12", 42},
				{"1\n3\n4\n2\n5\n7\n18\n100", 140},
			},
		},
		{
			title: "Input with new line and comma separated values should return the sum of all values.",
			testCases: []testCase{
				{"1\n3,4,2", 10},
				{"1\n3\n4,2", 10},
				{"1\n3,4\n2", 10},
				{"1,3\n4,2", 10},
			},
		},
	}

	for _, step := range steps {
		t.Run(fmt.Sprintf(step.title), func(t *testing.T) {
			for i, testCase := range step.testCases {
				t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {

					input := testCase.input
					expected := testCase.expected

					res, err := Sum(input)

					printableInput := strings.ReplaceAll(input, "\n", "\\n")

					if err != nil {
						t.Fatalf(`Sum(%v) returned an error %v. Was expecting no error and the value %d.`, printableInput, err, expected)
					}

					if res != expected {
						t.Fatalf(`Sum(%v) returned %d. Was expecting %d.`, printableInput, res, expected)
					}
				})
			}
		})
	}
}

func Test_Sum_CSV_ReturnsSumOfValues(t *testing.T) {
	input := "1,3,4,2,5,7,18,100"
	var expected int = 140

	res, err := Sum(input)

	if err != nil {
		t.Fatalf(`Sum(%v) returned an error %v. Was expecting no error and the value %d.`, input, err, expected)
	}

	if res != expected {
		t.Fatalf(`Sum(%v) returned %d. Was expecting %d.`, input, res, expected)
	}
}

func Test_Sum_NLSV_ReturnsSumOfValues(t *testing.T) {
	input := "1\n3\n4\n2"
	expected := 10

	res, err := Sum(input)

	if err != nil {
		t.Fatalf(`Sum(%v) returned an error %v. Was expecting no error and the value %d.`, input, err, expected)
	}

	if res != expected {
		t.Fatalf(`Sum(%v) returned %d. Was expecting %d.`, input, res, expected)
	}
}
