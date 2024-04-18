package interestCalc

import (
	"fmt"
	"math"
)

// CalculateCompoundInterest рассчитывает сложный процент для заданных параметров.
func CalculateCompoundInterest(principal, rate, time float64) float64 {
	amount := principal * math.Pow((1+rate/100), time)
	return amount - principal
}

// Message формирует сообщение о сложном проценте на основе переданных параметров.
// Принимает начальную сумму (principal), годовую ставку (rate) и количество лет (time).
// Возвращает строку с сообщением о сложном проценте.
func Message(interest, time float64) string {
	return fmt.Sprintf("Сложный процент за %.1f лет: %.2f", time, interest)
}

func NewPlus(interest, time float64) string {
	return fmt.Sprintf("Сложный процент за %.1f лет: %.2f", time, interest)
}
