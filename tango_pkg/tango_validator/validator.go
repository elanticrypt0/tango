package tango_validator

import (
	"fmt"
	"net/mail"
	"text/template"

	passwordstrenght "tango_pkg/passwordstrenghtchecker"
	"tango_pkg/tango_errors"
)

// Validate and satinize email
func ValidateEmail(email string) (string, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", tango_errors.ReturnDefault("Email", "Invalid email", 0)
	}
	return template.HTMLEscapeString(email), nil
}

// Validate and satinize password
func ValidatePassword(password string, minLen, maxLen int) (string, error) {

	passChecker := passwordstrenght.NewPasswordStrenght(minLen, maxLen)

	if password != "" {
		if !passChecker.HasMinLen(password) {
			return "", tango_errors.ReturnDefault("Password", fmt.Sprintf("Password must have %d chars minium", minLen), 0)
		}
		password = template.HTMLEscapeString(password)
		return password, nil
	} else {
		return "", &tango_errors.DefaultError{
			Name:    "Password",
			Code:    0,
			Message: "Invalid password",
		}
	}
}

// Validate and satinize password and verify is not weak
func ValidatePasswordStrong(password string, minLen, maxLen int) (string, error) {

	passChecker := passwordstrenght.NewPasswordStrenght(minLen, maxLen)

	if password != "" {
		passEval := passChecker.CheckAndEvaluate(password)
		if passEval.Strenght == "weak" {
			return "", tango_errors.ReturnDefault("Password", "Password must be stronger.", 0)
		}
		password = template.HTMLEscapeString(password)
		return password, nil
	} else {
		return "", &tango_errors.DefaultError{
			Name:    "Password",
			Code:    0,
			Message: "Invalid password",
		}
	}

}
