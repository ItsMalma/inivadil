package inivadil

import "reflect"

func Boolean() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() != reflect.Bool {
			return NewValidationError(p.Field, "field %v must be boolean", p.Field)
		}
		return NilValidationError()
	}
}

func String() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.Kind() != reflect.String {
			return NewValidationError(p.Field, "field %v must be string", p.Field)
		}
		return NilValidationError()
	}
}

func Integer() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.CanInt() {
			return NewValidationError(p.Field, "field %v must be string", p.Field)
		}
		return NilValidationError()
	}
}

func Float() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.CanFloat() {
			return NewValidationError(p.Field, "field %v must be string", p.Field)
		}
		return NilValidationError()
	}
}

func Array() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		switch p.RValue.Kind() {
		case reflect.Array, reflect.Slice:
			return NilValidationError()
		}
		return NewValidationError(p.Field, "field %v must be string", p.Field)
	}
}
