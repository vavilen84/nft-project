package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"time"
	"unicode/utf8"
)

func CustomFutureValidator(fl validator.FieldLevel) bool {
	if fl.Field().Type() != reflect.TypeOf(time.Time{}) {
		return false
	}
	timeValue, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	if timeValue.After(time.Now()) {
		return true
	}
	return false
}

func CustomPasswordValidator(fl validator.FieldLevel) bool {
	p := fl.Field().String()
	length := utf8.RuneCountInString(p)
	if length < 8 {
		return false
	}
	r, err := regexp.Match(alphaNumericRegexString, []byte(p))
	if err != nil {
		fmt.Println(err.Error())
	}
	return r
}
