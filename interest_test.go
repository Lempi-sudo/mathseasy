package interestCalc

import (
	"testing"
)

func TestCalculateCompoundInterest(t *testing.T) {
	// Тестовые данные
	principal := 1000.0
	rate := 5.0
	time := 2.0
	expected := 102.50 // Ожидаемый результат

	// Вызываем функцию для расчета сложного процента
	result := CalculateCompoundInterest(principal, rate, time)

	// Проверяем, что результат соответствует ожидаемому
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
