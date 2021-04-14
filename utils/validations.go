package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ValidatorUtil struct{}

func NewValidatorUtil() *ValidatorUtil {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
	return &ValidatorUtil{}
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func (ValidatorUtil) CreateMovieValidator(v validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)
	for _, f := range v {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err

	}
	return errs
}
