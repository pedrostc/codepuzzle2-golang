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
	input         string
	expected      int
	errorMessage  string
	errorExpected bool
}

func Test_challenge_steps(t *testing.T) {
	steps := []challengeStep{
		{
			title: "given an empty string should return 0.",
			testCases: []testCase{
				{input: "", expected: 0},
			},
		},
		{
			title: "given a string with a single value should return the parsed value",
			testCases: []testCase{
				{input: "1", expected: 1},
				{input: "2", expected: 2},
				{input: "3", expected: 3},
				{input: "42", expected: 42},
				{input: "100", expected: 100},
			},
		},
		{
			title: "given a comma separated string returns the sum of the values",
			testCases: []testCase{
				{input: "1,2,3,4", expected: 10},
				{input: "40,0,2", expected: 42},
				{input: "1,14,7,8,12", expected: 42},
				{input: "1,3,4,2,5,7,18,100", expected: 140},
			},
		},
		{
			title: "New line separated string returns the sum of the values",
			testCases: []testCase{
				{input: "1\n2\n3\n4", expected: 10},
				{input: "40\n0\n2", expected: 42},
				{input: "1\n14\n7\n8\n12", expected: 42},
				{input: "1\n3\n4\n2\n5\n7\n18\n100", expected: 140},
			},
		},
		{
			title: "Input with new line and comma separated values should return the sum of all values.",
			testCases: []testCase{
				{input: "1\n3,4,2", expected: 10},
				{input: "1\n3\n4,2", expected: 10},
				{input: "1\n3,4\n2", expected: 10},
				{input: "1,3\n4,2", expected: 10},
			},
		},
		{
			title: "should also support any 1 char user defined symbol delimiter using this format: '//(delimiter)\\n(numbers…)'.",
			testCases: []testCase{
				{input: "//;\n1;2;3", expected: 6},
				{input: "//-\n1-2-3-4-5", expected: 15},
				{input: "///\n4/6/3/7/1/1/1/1/8/1/9", expected: 42},
				{input: "//&\n1&1&1&1&1&1", expected: 6},
			},
		},
		{
			title: "should not accept negative numbers, throwing an error specifying the problematic numbers.",
			testCases: []testCase{
				{input: "1,-2", errorExpected: true, errorMessage: "negatives not allowed: -2"},
				{input: "-1\n-2,3,-4", errorExpected: true, errorMessage: "negatives not allowed: -1,-2,-4"},
				{input: "///\n-4/6/3/-7/1/-1/1/-1/8/1/9", errorExpected: true, errorMessage: "negatives not allowed: -4,-7,-1,-1"},
				{input: "//*\n-1*-2*-10", errorMessage: "negatives not allowed: -1,-2,-10"},
			},
		},
		{
			title: "should ignore (not add) numbers greater than 1000.",
			testCases: []testCase{
				{input: "1001,2", expected: 2},
				{input: "1000,2", expected: 1002},
				{input: "///\n2000/6/1/1/1234/5/3000/8/1/9", expected: 31},
				{input: "1\n2000,1\n10", expected: 12},
			},
		},
		{
			title: "should accept multi-character custom delimiter using this format: '//[delimiter]\\n(numbers…)'.",
			testCases: []testCase{
				{input: "//[;;;]\n1;;;2;;;3", expected: 6},
				{input: "//[-_-]\n1-_-2-_-3-_-4-_-5", expected: 15},
				{input: "//[//]\n4//6//3//7//1//1//1//1//8//1//9", expected: 42},
				{input: "//[&.?!]\n1&.?!1&.?!1&.?!1&.?!1&.?!1", expected: 6},
			},
		},
		{
			title: "should allow multiple single character delimiters using this format: '//[delim1][delim2]...\\n(numbers…)'.",
			testCases: []testCase{
				{input: "//[;][*]\n1;2*3", expected: 6},
				{input: "//[/][*]\n1/2/3**4/5", expected: 15},
				{input: "//[:][_][^][-]\n4:6-3-7_1-1^1:1_8^1^9", expected: 42},
				{input: "//[+][*][^]\n1^1+1^1*1*1", expected: 6},
			},
		},
		{
			title: "should allow multiple multi-characters delimiters using this format: '//[delim1][delim2]...\\n(numbers…)'.",
			testCases: []testCase{
				{input: "//[**][;]\n1;2**3", expected: 6},
				{input: "//[//][***]\n1//2//3***4//5", expected: 15},
				{input: "//[:)][:(]\n4:)6:(3:(7:)1:)1:)1:(1:)8:)1:)9", expected: 42},
			},
		},
	}

	for _, step := range steps {
		t.Run(fmt.Sprintf(step.title), func(t *testing.T) {
			for i, testCase := range step.testCases {
				t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
					if !testCase.errorExpected {
						executeTest(testCase, t)
					} else {
						executeErrorTest(testCase, t)
					}
				})
			}
		})
	}
}

func executeTest(testCase testCase, t *testing.T) {
	if testCase.errorExpected {
		t.Fatalf("The test case is running on the wrong executor.\nPlease check the test case configuration.")
	}

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
}

func executeErrorTest(testCase testCase, t *testing.T) {
	if !testCase.errorExpected {
		t.Fatalf("The test case is running on the wrong executor.\nPlease check the test case configuration.")
	}

	input := testCase.input
	expected := testCase.errorMessage

	res, err := Sum(input)

	printableInput := strings.ReplaceAll(input, "\n", "\\n")

	if err == nil {
		t.Fatalf(`Sum(%v) did not returned the value %d. Was expecting an error with the message "%v". `, printableInput, res, expected)
	}

	actual := err.Error()

	if err.Error() != expected {
		t.Fatalf(`Sum(%v) returned the error "%v". Was expecting "%v".`, printableInput, actual, expected)
	}
}
