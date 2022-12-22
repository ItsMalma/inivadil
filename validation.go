package inivadil

import (
	"encoding/json"
	"io"
	"reflect"
)

type ValidationType uint8

const (
	JSONValidation ValidationType = iota
)

func Validate(t ValidationType, rules Rules, r io.Reader) error {
	switch t {
	case JSONValidation:
		return validateJSON(rules, r)
	}
	return ErrInvalidValidationType
}

func validateJSON(rules Rules, r io.Reader) error {
	m := map[string]any{}
	if err := json.NewDecoder(r).Decode(&m); err != nil {
		return err
	}
	validationErrors := ValidationErrors{}
	for field, rule := range rules {
		value := m[field]
	elementLoop:
		for _, elem := range rule {
			param := RuleElementParameter{
				Field:  field,
				Value:  value,
				RValue: reflect.ValueOf(value),
			}
			err := elem(param)
			if err.stop {
				break elementLoop
			}
			if !err.Nil() {
				validationErrors = append(validationErrors, err)
			}
		}
	}
	if len(validationErrors) < 1 {
		return nil
	}
	return validationErrors
}
