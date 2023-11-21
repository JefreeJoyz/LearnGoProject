package main

import "testing"

func TestMax(t *testing.T) {
	// arrange
	numbers := []int{1, 5, 22, 11, 123}
	expected := 123
	// act
	result := Max(numbers)
	t.Logf("calling %v, result is %v", numbers, result) // Это если нам надо глянуть чуть детальнее
	// assert
	if result != expected {
		t.Errorf("Incorrect result. Expected %d, result %d", expected, result)
	}
}
