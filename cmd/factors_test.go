package cmd

import (
  "strconv" 
  "testing" 
)

func TestIsPrime(t *testing.T) {
  testCases:=[]struct{
    input int
    output bool
  } {
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
    t.Run(strconv.Itoa( testCase.input ), func(t *testing.T) {
      ans:=isPrime(testCase.input)
      if ans!=testCase.output {
        t.Errorf("%d is prime, expected %t, got %t", testCase.input, testCase.output, ans)
      }
    })
  }
}


func TestGetFactors(t *testing.T) {
  testCases:=[]struct{
    input int
    output []int
  } {
    {1, []int{1}},
    {2, []int{2}},
    {3, []int{3}},
    {4, []int{2,2}},
    {35, []int{5,7}},
    {210, []int{2,3,5,7}},
    {1050, []int{2,3,5,5,7}},
  }
  for _, testCase := range testCases {
    t.Run(strconv.Itoa( testCase.input ), func(t *testing.T) {
      ans:=getPrimeFactors(testCase.input)
      if !Equal(ans, testCase.output) {
        t.Errorf("factors for %d, expected %v, got %v", testCase.input, testCase.output, ans)
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
