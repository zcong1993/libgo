package validator

import "gopkg.in/go-playground/validator.v9"

func NormalizeErr(err error) validator.ValidationErrorsTranslations {
	if err == nil {
		return nil
	}

	errs, ok := err.(validator.ValidationErrors)

	if !ok {
		return nil
	}

	return errs.Translate(trans)
}
