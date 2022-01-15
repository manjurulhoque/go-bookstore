package utils

import (
	"encoding/json"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
	"strings"
)

type IError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func TranslateError(err error, trans ut.Translator) (errs []IError) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		//translatedErr := fmt.Errorf(e.Translate(trans))
		translatedErr := IError{
			Field:   strings.ToLower(e.Field()),
			Message: e.Translate(trans),
		}
		errs = append(errs, translatedErr)
	}
	return errs
}
