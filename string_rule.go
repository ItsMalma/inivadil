package inivadil

import (
	"encoding/base32"
	"encoding/base64"
	"net/mail"
	"reflect"
	"strconv"
	"strings"
)

func BooleanString() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if (p.RValue.Kind() == reflect.String) && (p.Value == "true" || p.Value == "false") {
			return NilValidationError()
		}
		return NewValidationError(p.Field, "field %v must be boolean as a string", p.Field)
	}
}

func IntegerString() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() == reflect.String {
			if _, err := strconv.ParseInt(p.RValue.String(), 10, 64); err == nil {
				return NilValidationError()
			}
		}
		return NewValidationError(p.Field, "field %v must be number as a string", p.Field)
	}
}

func FloatString() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() == reflect.String {
			if _, err := strconv.ParseFloat(p.RValue.String(), 64); err == nil {
				return NilValidationError()
			}
		}
		return NewValidationError(p.Field, "field %v must be number as a string", p.Field)
	}
}

func Contains(str string) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() == reflect.String {
			if strings.Contains(p.RValue.String(), str) {
				return NilValidationError()
			}
		}
		return NewValidationError(p.Field, "field %v must be a string that contains %v", p.Field, str)
	}
}

func NotContains(str string) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() == reflect.String {
			if !strings.Contains(p.RValue.String(), str) {
				return NilValidationError()
			}
		}
		return NewValidationError(p.Field, "field %v mustn't be a string that contains %v", p.Field, str)
	}
}

func OnlyAlpha() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() == reflect.String {
			stringValue := p.RValue.String()
			if strings.Map(func(r rune) rune {
				if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
					return r
				}
				return -1
			}, stringValue) == stringValue {
				return NilValidationError()
			}
		}
		return NewValidationError(p.Field, "field %v must be a string that only contains alphabetical characters", p.Field)
	}
}

func OnlyAlphanumeric() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() == reflect.String {
			stringValue := p.RValue.String()
			if strings.Map(func(r rune) rune {
				if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
					return r
				}
				return -1
			}, stringValue) == stringValue {
				return NilValidationError()
			}
		}
		return NewValidationError(p.Field, "field %v must be a string that only contains alphabetical and numerical characters", p.Field)
	}
}

func OnlyASCII() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() == reflect.String {
			stringValue := p.RValue.String()
			if strings.Map(func(r rune) rune {
				if r <= 127 {
					return r
				}
				return -1
			}, stringValue) == stringValue {
				return NilValidationError()
			}
		}
		return NewValidationError(p.Field, "field %v must be a string that only contains ascii characters", p.Field)
	}
}

func Base32() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() == reflect.String {
			if _, err := base32.StdEncoding.DecodeString(p.RValue.String()); err == nil {
				return NilValidationError()
			}
		}
		return NewValidationError(p.Field, "field %v must be a string that base32 encoded", p.Field)
	}
}

func Base64() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() == reflect.String {
			if _, err := base64.StdEncoding.DecodeString(p.RValue.String()); err == nil {
				return NilValidationError()
			}
		}
		return NewValidationError(p.Field, "field %v must be a string that base32 encoded", p.Field)
	}
}

func Email() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() == reflect.String {
			if _, err := mail.ParseAddress(p.RValue.String()); err == nil {
				return NilValidationError()
			}
		}
		return NewValidationError(p.Field, "field %v must be a string that base32 encoded", p.Field)
	}
}
