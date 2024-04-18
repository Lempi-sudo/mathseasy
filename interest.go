package complex_interest

import (
	"math"
)

// CalculateCompoundInterest рассчитывает сложный процент для заданных параметров.
func CalculateCompoundInterest(principal float64, rate float64, time float64) float64 {
	amount := principal * math.Pow((1+rate/100), time)
	return amount - principal
}
