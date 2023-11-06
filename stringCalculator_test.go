package stringCalculator

import "testing"

func Test_Sum_EmptyString_Returns(t *testing.T) {
	input := ""
	expected := 0

	res, err := Sum(input)

	if err != nil {
		t.Fatalf(`Sum(%v) returned an error %v. Was expecting no error and the value %d.`, input, err, expected)
	}

	if res != expected {
		t.Fatalf(`Sum(%v) returned %d. Was expecting %d.`, input, res, expected)
	}
}

func Test_Sum_SingleValue_ReturnsParsedValue(t *testing.T) {
	input := "2"
	var expected int = 2

	res, err := Sum(input)

	if err != nil {
		t.Fatalf(`Sum(%v) returned an error %v. Was expecting no error and the value %d.`, input, err, expected)
	}

	if res != expected {
		t.Fatalf(`Sum(%v) returned %d. Was expecting %d.`, input, res, expected)
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

func Test_Sum_NLSV_and_CSV_ReturnsSumOfValues(t *testing.T) {
	input := "1\n3,4,2"
	expected := 10

	res, err := Sum(input)

	if err != nil {
		t.Fatalf(`Sum(%v) returned an error %v. Was expecting no error and the value %d.`, input, err, expected)
	}

	if res != expected {
		t.Fatalf(`Sum(%v) returned %d. Was expecting %d.`, input, res, expected)
	}
}
