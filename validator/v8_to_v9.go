package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

// DefaultValidator impl the gin binding StructValidator
// use binding.Validator = new(DefaultValidator)
type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
	trans    ut.Translator
}

var _ binding.StructValidator = &DefaultValidator{}

// ValidateStruct impl binding.StructValidator's ValidateStruct
func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

// Engine impl binding.StructValidator's Engine
func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
		enLang := en.New()
		uni = ut.New(enLang, enLang)

		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		trans, _ = uni.GetTranslator("en")
		v.trans = trans

		en_translations.RegisterDefaultTranslations(v.validate, trans)
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
