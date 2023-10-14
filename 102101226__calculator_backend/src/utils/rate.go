package utils

import (
	"fmt"
)

func StoreRateToMoney(money, rate float64) (string, error) {
	interest := money * rate
	return fmt.Sprint(interest), nil
}

func ProvideRateToMoney(money, rate float64) (string, error) {
	interest := money * rate
	return fmt.Sprint(interest), nil
}
