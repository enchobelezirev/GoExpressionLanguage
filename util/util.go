package util

import "strconv"

func ConvertTokenToInteger(token string) (int, error) {
	return strconv.Atoi(token)
}

func IsOperand(token string) bool {
	_, err := ConvertTokenToInteger(token)
	if err != nil {
		return false
	}
	return true
}
