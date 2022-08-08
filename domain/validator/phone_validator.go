package validator

import (
	"errors"
	"phone-script/domain"
	"strconv"
)

func Validate(phone domain.Phone) error {
	_, err := strconv.Atoi(phone.Number)
	if err != nil {
		return errors.New("invalid number")
	}

	numLength := len(phone.Number)
	if (numLength > 12 || numLength < 7) && numLength != 3 {
		return errors.New("number doesn't match")
	}

	return nil
}
