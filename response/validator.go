package response

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	t_en "github.com/go-playground/validator/v10/translations/en"
	t_zh "github.com/go-playground/validator/v10/translations/zh"
)

var Translator ut.Translator

func InitValidatorTranslator(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		_zh := zh.New() // 中文翻译器
		_en := en.New() // 英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是支持的语言环境
		Translator, ok = ut.New(_en, _zh, _en).GetTranslator(local)
		if !ok {
			return fmt.Errorf("GetTranslator(%s)", local)
		}

		switch local {
		case "en":
			err = t_en.RegisterDefaultTranslations(v, Translator)
		case "zh":
			err = t_zh.RegisterDefaultTranslations(v, Translator)
		default:
			err = t_en.RegisterDefaultTranslations(v, Translator)
		}
		return
	}
	return
}
