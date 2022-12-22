package inivadil

import "reflect"

func NotNull() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.Value == nil {
			return NewValidationError(p.Field, "field %v must be not null", p.Field)
		}
		return NilValidationError()
	}
}

func Equals(value any) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.Value != value {
			return NewValidationError(p.Field, "field %v must be %v", p.Field, value)
		}
		return NilValidationError()
	}
}

func NotEquals(value any) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.Value == value {
			return NewValidationError(p.Field, "field %v mustn't be %v", p.Field, value)
		}
		return NilValidationError()
	}
}

func Empty() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		switch p.RValue.Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
			if p.RValue.Len() != 0 {
				return NewValidationError(p.Field, "field %v must be empty", p.Field)
			}
		case reflect.Pointer:
			if p.RValue.Elem().Kind() == reflect.Array && p.RValue.Len() != 0 {
				return NewValidationError(p.Field, "field %v must be empty", p.Field)
			}
		}
		return NilValidationError()
	}
}

func NotEmpty() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		switch p.RValue.Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
			if p.RValue.Len() == 0 {
				return NewValidationError(p.Field, "field %v mustn't be empty", p.Field)
			}
		case reflect.Pointer:
			if p.RValue.Elem().Kind() == reflect.Array && p.RValue.Len() == 0 {
				return NewValidationError(p.Field, "field %v mustn't be empty", p.Field)
			}
		}
		return NilValidationError()
	}
}

func Length(min, max int) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		switch p.RValue.Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
			len := p.RValue.Len()
			if len < min || len > max {
				return NewValidationError(p.Field, "field %v's length must be greater than or equal to %v and less than or equal to %v", p.Field, min, max)
			}
		case reflect.Pointer:
			if p.RValue.Elem().Kind() == reflect.Array {
				len := p.RValue.Len()
				if len < min || len > max {
					return NewValidationError(p.Field, "field %v's length must be greater than or equal to %v and less than or equal to %v", p.Field, min, max)
				}
			}
		}
		return NilValidationError()
	}
}

func MinLength(min int) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		switch p.RValue.Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
			if p.RValue.Len() < min {
				return NewValidationError(p.Field, "field %v's length must be greater than or equal to %v", p.Field, min)
			}
		case reflect.Pointer:
			if p.RValue.Elem().Kind() == reflect.Array && p.RValue.Len() >= min {
				return NewValidationError(p.Field, "field %v's length must be greater than or equal to %v", p.Field, min)
			}
		}
		return NilValidationError()
	}
}

func MaxLength(max int) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		switch p.RValue.Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
			if p.RValue.Len() > max {
				return NewValidationError(p.Field, "field %v's length must be less than or equal to %v", p.Field, max)
			}
		case reflect.Pointer:
			if p.RValue.Elem().Kind() == reflect.Array && p.RValue.Len() >= max {
				return NewValidationError(p.Field, "field %v's length must be less than or equal to %v", p.Field, max)
			}
		}
		return NilValidationError()
	}
}
