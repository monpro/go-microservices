package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translation() gin.HandlerFunc {
	return func(context *gin.Context) {
		uniLanguage := ut.New(en.New(), zh.New())
		locale := context.GetHeader("locale")
		translation, _ := uniLanguage.GetTranslator(locale)
		value, status := binding.Validator.Engine().(*validator.Validate)
		if status {
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(value, translation)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(value, translation)
				break
			default:
				_ = en_translations.RegisterDefaultTranslations(value, translation)
				break
			}
			context.Set("trans", translation)
		}
		context.Next()
	}
}