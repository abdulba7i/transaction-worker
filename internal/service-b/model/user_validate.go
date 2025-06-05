package model

import (
	"fmt"
)

func (u *User) ValidateInput() error {
	if u.UserID <= 0 {
		return fmt.Errorf("UserID не может быть меньше нуля")
	}

	if u.RequestID == "" {
		return fmt.Errorf("RequestID не может быть пустым")
	}

	if u.Amount <= 0 || u.Amount > 1000 {
		return fmt.Errorf("Введите значение от 0 до 1000")
	}

	return nil
}

// эта функция должна быть в сервисе бизнес-логики
// func CheckDuplicateRequest() {

// }
