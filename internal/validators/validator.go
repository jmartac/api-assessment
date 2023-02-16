package validators

import (
	"github.com/go-playground/validator/v10"
	re "regexp"
)

type Validator struct {
	validator *validator.Validate
}

// NewValidator is used to create a new validator
func NewValidator() *Validator {
	validate := validator.New()
	err := validate.RegisterValidation("username", usernameValidator)
	if err != nil {
		return nil
	}
	err = validate.RegisterValidation("password", passwordValidator)
	if err != nil {
		return nil
	}
	return &Validator{
		validator: validate,
	}
}

// Validate is used to validate a struct
// If the struct has no validation tags, it will always return nil
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// usernameValidator is used to validate the username field of a user request
// The username must start with a letter, it can only contain letters and numbers and be between 3 and 30 characters long
func usernameValidator(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	reg := re.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{2,29}$`)
	return reg.MatchString(username)
}

// passwordValidator is used to validate the password field of a user request
// The password must start with a letter, it can only contain letters and numbers and be between 8 and 128 characters long
func passwordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	reg := re.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{7,127}$`)
	return reg.MatchString(password)
}
