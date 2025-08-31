package libuser

import (
	"context"
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/yeencloud/lib-shared/validation"
)

var usernameRegexp = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]+$")
var validationCodeRegexp = regexp.MustCompile("^[0-9]+$")

func userNameValidator(ctx context.Context, fl validator.FieldLevel) error {
	username := fl.Field().String()

	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}

	if !usernameRegexp.MatchString(username) {
		return errors.New("username must start with a letter and can only contain numbers and letters")
	}

	return nil
}

func validationCodeValidator(ctx context.Context, fl validator.FieldLevel) error {
	validationCode := fl.Field().String()

	if len(validationCode) != 6 {
		return errors.New("validation code must be 6 characters long")
	}

	if !validationCodeRegexp.MatchString(validationCode) {
		return errors.New("validation code can only contain digits")
	}

	return nil
}

func Validations() map[string]validation.ValidationFunc {
	return map[string]validation.ValidationFunc{
		"username":        userNameValidator,
		"validation_code": validationCodeValidator,
	}
}
