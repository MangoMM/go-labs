package calculator

type Calculator struct {}

func (c Calculator) Sum(x... float64) float64 {
	sum := 0.0

	for _, v := range x {
		sum += v
	}

	return sum
}

func (c Calculator) Multiply(x... float64) float64 {
	ans := 0.00

	for i, v := range x {
		if i == 0 {
			ans = v
			continue
		}

		ans *= v
	}
	return ans
}

func (c Calculator) Average(x... float64) float64 {
	return 0.0
}
