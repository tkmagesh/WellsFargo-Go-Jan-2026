package utils

import "testing"

func Test_IsPrime_79(t *testing.T) {
	t.Skip("skipping Test_IsPrime_79")
	// arrange
	sut := 79
	expectedResult := true

	// act
	actualResult := IsPrime(sut)

	// assert
	if actualResult != expectedResult {
		/*
			t.Logf("IsPrime(79), expected = %t, actual = %t\n", expectedResult, actualResult)
			t.Fail()
		*/
		t.Errorf("IsPrime(79), expected = %t, actual = %t\n", expectedResult, actualResult)
	}
}

// data driven test
func Test_IsPrime(t *testing.T) {
	var testData = []struct {
		name     string
		no       int
		expected bool
	}{
		{name: "IsPrime-[13]", no: 13, expected: true},
		{name: "IsPrime-[15]", no: 15, expected: false},
		{name: "IsPrime-[17]", no: 17, expected: true},
		{name: "IsPrime-[21]", no: 21, expected: false},
		{name: "IsPrime-[79]", no: 79, expected: true},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			actual := IsPrime(td.no)
			if actual != td.expected {
				t.Errorf("%q, expected = %t, actual = %t\n", td.name, td.expected, actual)
			}
		})
	}
}

// Benchmark
func Benchmark_GeneratePrimes(b *testing.B) {
	// setup
	for b.Loop() {
		GeneratePrimes(2, 100)
	}
	// cleanup
}
