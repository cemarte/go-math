package lib

import (
	"fmt"
	"strconv"
	"testing"
)

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		input  int
		output bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
		{17, true},
		{31, true},
		{900003043, true},
	}

	for _, testCase := range testCases {
		t.Run(strconv.Itoa(testCase.input), func(t *testing.T) {
			ans := IsPrime(testCase.input)
			if ans != testCase.output {
				t.Errorf("%d is prime, expected %t, got %t", testCase.input, testCase.output, ans)
			}
		})
	}
}

func TestGetFactors(t *testing.T) {
	testCases := []struct {
		input  int
		output []int
	}{
		{1, []int{1}},
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{35, []int{5, 7}},
		{210, []int{2, 3, 5, 7}},
		{1050, []int{2, 3, 5, 5, 7}},
	}
	for _, testCase := range testCases {
		t.Run(strconv.Itoa(testCase.input), func(t *testing.T) {
			ans := GetPrimeFactors(testCase.input)
			if !Equal(ans, testCase.output) {
				t.Errorf("factors for %d, expected %v, got %v", testCase.input, testCase.output, ans)
			}
		})
	}
}

func TestGetMedian(t *testing.T) {
	testCases := []struct {
		input  []int
		output float32
	}{
		{[]int{}, 0.0},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
		{[]int{0, 0, 3, 6, 6}, 3},
		{[]int{3, 1, 6, 9, 0}, 3},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.input), func(t *testing.T) {
			ans := GetMedian(testCase.input)
			if ans != testCase.output {
				t.Errorf("Expected %f, got %f", testCase.output, ans)
			}
		})
	}
}

type modeResult struct {
	values  []int
	hasMode bool
}

type modeTestCase struct {
	input  []int
	output modeResult
}

func TestGetMode(t *testing.T) {

	testCases := []modeTestCase{
		{[]int{}, modeResult{[]int{}, false}},
		{[]int{1, 2, 3, 4, 5, 6}, modeResult{[]int{}, false}},
		{[]int{0, 0, 3, 6, 6}, modeResult{[]int{0, 6}, true}},
		{[]int{3, 1, 6, 9, 0, 3}, modeResult{[]int{3}, true}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.input), func(t *testing.T) {
			ans, hasMode := GetMode(testCase.input)
			if !Equal(ans, testCase.output.values) || hasMode != testCase.output.hasMode {
				t.Errorf("Expected %v, got %v", testCase.output.values, ans)
			}
		})
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
