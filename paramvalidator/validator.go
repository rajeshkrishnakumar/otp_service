package paramvalidator

import (
	"errors"
	"regexp"

	validator "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func IsMobileNumber(mobile string) error {
	if len(mobile) <= 6 || len(mobile) >= 15 {
		return errors.New("Mobile number invalid length")
	}
	r := regexp.MustCompile(`\d{10}`)
	if !r.MatchString(mobile) {
		return errors.New("Mobile number is invalid")
	}
	return nil
}

func IsEmail(email string) error {
	validate = validator.New()
	errs := validate.Var(email, "email")
	return errs
}
