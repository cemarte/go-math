package lib

import (
	"sort"
)

var knownPrimes = map[int]struct{}{
	2:  struct{}{},
	3:  struct{}{},
	5:  struct{}{},
	7:  struct{}{},
	11: struct{}{},
	13: struct{}{},
	17: struct{}{},
	19: struct{}{},
	23: struct{}{},
	29: struct{}{},
}

func IsPrime(i int) bool {
	if i <= 1 {
		return false
	}

	if _, ok := knownPrimes[i]; ok {
		return true
	}

	for j := 2; j <= i/2; j++ {
		if i%j == 0 {
			return false
		}
	}

	knownPrimes[i] = struct{}{}
	return true
}

func GetPrimeFactors(i int) []int {
	if i == 0 || i == 1 || i == 2 || i == 3 {
		return []int{i}
	}

	result := []int{}

	for v := range knownPrimes {
		for i%v == 0 {
			result = append(result, v)
			i = i / v
		}
		if i == 1 {
			break
		}
	}

	// TODO: ordering isn't the same between test runs, check why, I think it relates to array capacity
	sort.Ints(result)
	return result
}

func GetMedian(input []int) float32 {

	length := len(input)
	if length == 0 {
		return float32(0)
	}

	sort.Ints(input)

	if length%2 == 0 {
		return (float32(input[(length-1)/2]) + float32(input[(length/2)])) / 2.0
	} else {
		return float32(input[length/2])
	}
}

func GetMode(input []int) ([]int, bool) {
	hash := make(map[int]int)

	for _, val := range input {
		hash[val] += 1
	}

	maxCount := 0
	results := make([]int, 0)

	for number, occurrences := range hash {
		if occurrences == maxCount {
			results = append(results, number)
		} else if occurrences > maxCount {
			results = []int{number}
			maxCount = occurrences
		}
	}

	if len(results) == len(input) {
		return []int{}, false
	}

	return results, maxCount > 0
}
