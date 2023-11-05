package stringCalculator

import "testing"

func Test_Sum_EmptyString_Returns(t *testing.T) {
	res, err := Sum("")
	expected := 0

	if res != expected || err != nil {
		t.Fatalf(`Sum("") = %d, %v, want match for %d, nil`, res, err, expected)
	}
}

func Test_Sum_SingleValue_ReturnsParsedValue(t *testing.T) {
	res, err := Sum("1")
	var expected int = 2

	if res != expected || err != nil {
		t.Fatalf(`Sum("") = %d, %v, want match for %d, nil`, res, err, expected)
	}
}

func Test_Sum_CSV_ReturnsSumOfValues(t *testing.T) {
	res, err := Sum("1,3,4,2,5,7,18,100")
	var expected int = 140
	if res != expected || err != nil {
		t.Fatalf(`Sum("") = %d, %v, want match for %d, nil`, res, err, expected)
	}
}

func Test_Sum_NLSV_ReturnsSumOfValues(t *testing.T) {
	res, err := Sum("1\n3\n4\n2")
	var expected int = 10
	if res != expected || err != nil {
		t.Fatalf(`Sum("") = %d, %v, want match for %d, nil`, res, err, expected)
	}
}
