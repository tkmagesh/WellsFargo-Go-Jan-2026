package utils

func IsPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimes(start, end int) []int {
	var result []int = make([]int, 0, 40)
	for no := start; no <= end; no++ {
		if IsPrime(no) {
			result = append(result, no)
		}
	}
	return result
}
