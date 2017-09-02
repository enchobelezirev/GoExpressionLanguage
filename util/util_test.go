package util

import (
	"testing"
)

func TestConvertTokenToIntegerNoError(t *testing.T) {
	num, err := ConvertTokenToInteger("4567")
	if err != nil {
		t.Error("Did not expect error")
	}
	if num != 4567 {
		t.Errorf("Expected %d, got %d", 4567, num)
	}
}

func TestConvertTokenToIntegerWithError(t *testing.T) {
	_, err := ConvertTokenToInteger("drun-drun")
	if err == nil {
		t.Error("Expected ERROR")
	}
	expectedMessage := "strconv.Atoi: parsing \"drun-drun\": invalid syntax"
	if expectedMessage != err.Error() {
		t.Errorf("Expected %s, got %s", expectedMessage, err.Error())
	}
}

func TestIsOperandWithError(t *testing.T) {
	res := IsOperand("drun-drun")
	if res {
		t.Error("Expected to return false")
	}
}

func TestIsOperand(t *testing.T) {
	res := IsOperand("3464")
	if !res {
		t.Error("Expected to return false")
	}
}
