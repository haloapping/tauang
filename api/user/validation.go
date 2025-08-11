package user

import (
	"net/mail"
	"unicode"

	"github.com/nyaruka/phonenumbers"
)

func registerValidation(req registerReq) map[string][]string {
	errValidation := make(map[string][]string, 5)

	if req.Username == "" {
		errValidation["username"] = append(errValidation["username"], "cannot empty")
	}

	if req.Email == "" {
		errValidation["email"] = append(errValidation["email"], "cannot empty")
	}
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		errValidation["email"] = append(errValidation["email"], "format invalid")
	}

	if req.Phone == "" {
		errValidation["phone"] = append(errValidation["phone"], "cannot empty")
	}
	_, err = phonenumbers.Parse(req.Phone, "ID")
	if err != nil {
		errValidation["phone"] = append(errValidation["phone"], "format invalid")
	}

	nLowercase := 0
	nUppercase := 0
	nDigit := 0
	nPunct := 0
	for _, r := range req.Password {
		if unicode.IsLower(r) {
			nLowercase++
		}

		if unicode.IsUpper(r) {
			nUppercase++
		}

		if unicode.IsDigit(r) {
			nDigit++
		}

		if unicode.IsPunct(r) {
			nPunct++
		}
	}
	if len(req.Password) < 8 {
		errValidation["password"] = append(errValidation["password"], "min 8 chars")
	}
	if nLowercase < 1 {
		errValidation["password"] = append(errValidation["password"], "min 1 lowercase")
	}
	if nUppercase < 1 {
		errValidation["password"] = append(errValidation["password"], "min 1 uppercase")
	}
	if nDigit < 1 {
		errValidation["password"] = append(errValidation["password"], "min 1 digit")
	}
	if nPunct < 1 {
		errValidation["password"] = append(errValidation["password"], "min 1 punctuation")
	}

	if req.Password != req.ConfirmPassword {
		errValidation["confirmPassword"] = append(errValidation["confirmPassword"], "not match with password")
	}

	return errValidation
}

func loginValidation(req loginReq) map[string][]string {
	errValidation := make(map[string][]string, 5)

	if req.Username == "" {
		errValidation["username"] = append(errValidation["username"], "cannot empty")
	}

	if req.Password == "" {
		errValidation["password"] = append(errValidation["password"], "cannot empty")
	}

	return errValidation
}
