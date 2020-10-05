package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValid(context *gin.Context, value interface{}) (bool, ValidErrors)  {
	var errs ValidErrors
	err := context.ShouldBind(value)
	if err != nil {
		value := context.Value("trans")
		translation, _ := value.(ut.Translator)
		validation, status := err.(val.ValidationErrors)
		if !status {
			return false, errs
		}
		for key, value := range validation.Translate(translation) {
			errs = append(errs, &ValidError{
				Key: 	 key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}
