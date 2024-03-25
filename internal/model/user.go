package model

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/achsanit/my-gram/pkg/helper"
)

type UserRegister struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	DOB      time.Time `json:"dob"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	DOB       time.Time `json:"dob"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (u UserRegister) ValidateInput() error {

	errVal := map[string]interface{}{}

	if u.Username == "" {
		errVal["username"] = "username cant empty"
	}
	if len(u.Password) < 6 {
		errVal["password"] = "password must have min. 6 character"
	}
	if u.Email == "" || !helper.IsValidEmail(u.Email) {
		errVal["email"] = "invalid email format"
	}

	jsonErr, err := json.Marshal(errVal)

	if err != nil {
		return err
	}

	if len(errVal) > 0 {
		return errors.New(string(jsonErr))
	}

	return nil
}
