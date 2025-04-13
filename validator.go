package lib_user

import (
	"context"
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/yeencloud/lib-shared/validation"
)

func userNameValidator(ctx context.Context, fl validator.FieldLevel) error {
	username := fl.Field().String()

	if len(username) < 3 || len(username) > 20 {
		return errors.New("username must be between 3 and 20 characters long")
	}

	pattern := "^[a-zA-Z0-9_.-]+$"
	re := regexp.MustCompile(pattern) // TODO: MustCompile should be used only once for performance

	if !re.MatchString(username) {
		return errors.New("username can only contain letters, digits, underscores, hyphens, or dots")
	}

	return nil
}

func validationCodeValidator(ctx context.Context, fl validator.FieldLevel) error {
	validationCode := fl.Field().String()

	if len(validationCode) != 6 {
		return errors.New("validation code must be 6 characters long")
	}

	pattern := "^[0-9]+$"
	re := regexp.MustCompile(pattern) // TODO: MustCompile should be used only once for performance

	if !re.MatchString(validationCode) {
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
