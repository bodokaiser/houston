// Package model provides entities for safe device configuration.
//
// By using an extra model layer between the device drivers and network
// interfaces like HTTP we can control what information we want to expose
// to the outside world. On the one hand side we can catch invalid device
// configurations at an early stage on the other side we can hide
// (complicated) implementation details from the API user.
package model

import (
	"reflect"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("range", ValidateRange)
	validate.RegisterStructValidation(DDSParamValidation, DDSParam{})
}

// ValidateRange checks if an array of size two is a valid range.
//
// This means that the last value should be greater than the first value.
func ValidateRange(f validator.FieldLevel) bool {
	switch f.Field().Kind() {
	case reflect.Array:
	case reflect.Slice:
	default:
		return false
	}

	a, ok := f.Field().Index(0).Interface().(float64)
	if !ok {
		return false
	}
	// we skip the additional ok because we dont have mixed type arrays, slices
	b := f.Field().Index(f.Field().Len() - 1).Interface().(float64)

	return a < b
}

// DDSParamValidation implements struct level validation.
func DDSParamValidation(sl validator.StructLevel) {
	p := sl.Current().Interface().(DDSParam)
	i := 0

	if p.DDSConst != nil {
		i++
	}
	if p.DDSSweep != nil {
		i++
	}
	if p.DDSPlayback != nil {
		i++
	}

	if i == 0 {
		sl.ReportError(p.DDSConst, "DDSParam", "param", "noparam", "")
	}
	if i > 1 {
		sl.ReportError(p.DDSConst, "DDSParam", "param", "toomanyparam", "")
	}
}
