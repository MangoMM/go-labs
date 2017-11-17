package calculator

import (
	"testing"
	"fmt"
)

func TestCalculator_Sum(t *testing.T) {

	calc := Calculator{}

	sum := calc.Sum(1, 3, 7)

	if sum != 11.0 {
		t.Errorf("expected %0.2f, got %0.2f", 11.0, sum)
	}

	sum = calc.Sum(3)

	if sum != 3.0 {
		t.Errorf("expected %0.2f, got %0.2f", 3.0, sum)
	}
}

func TestCalculator_Multiply(t *testing.T) {

	type testPair struct {
		values []float64
		answer float64
	}

	testCases := []testPair{
		{answer: 21.0, values:[]float64{1,3,7}},
		{answer: 0.0, values:[]float64{}},
	}

	calc := Calculator{}
	for _, tc := range testCases {
		//ans := calc.Multiply(tc.values...)
		//if ans != tc.answer {
		//	t.Errorf("expected %0.2f, got %0.2f", tc.answer, ans)
		//}

		if ans := calc.Multiply(tc.values...); ans != tc.answer {
			t.Errorf("expected %0.2f, got %0.2f", tc.answer, ans)
		}
	}
}

func BenchmarkSprint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprint("HEllo World")
	}
}

func BenchmarkSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprintf("HEllo World")
	}
}
