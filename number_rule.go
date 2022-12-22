package inivadil

import "math"

func DivisibleBy(number float64) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.CanInt() {
			valueInt := p.RValue.Int()
			if math.Mod(float64(valueInt), number) == 0 {
				return NewValidationError(p.Field, "field %v must be divisible by %v", p.Field, number)
			}
		} else if p.RValue.CanFloat() {
			valueFloat := p.RValue.Float()
			if math.Mod(valueFloat, number) == 0 {
				return NewValidationError(p.Field, "field %v must be divisible by %v", p.Field, number)
			}
		}
		return NilValidationError()
	}
}

func Positive() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.CanInt() {
			valueInt := p.RValue.Int()
			if valueInt < 0 {
				return NewValidationError(p.Field, "field %v must be positive", p.Field)
			}
		} else if p.RValue.CanFloat() {
			valueFloat := p.RValue.Float()
			if valueFloat < 0 {
				return NewValidationError(p.Field, "field %v must be positive", p.Field)
			}
		}
		return NilValidationError()
	}
}

func Negative() RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.CanInt() {
			valueInt := p.RValue.Int()
			if valueInt > 0 {
				return NewValidationError(p.Field, "field %v must be negative", p.Field)
			}
		} else if p.RValue.CanFloat() {
			valueFloat := p.RValue.Float()
			if valueFloat > 0 {
				return NewValidationError(p.Field, "field %v must be negative", p.Field)
			}
		}
		return NilValidationError()
	}
}

func Min(minNumber float64) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.CanInt() {
			valueInt := p.RValue.Int()
			if float64(valueInt) < minNumber {
				return NewValidationError(p.Field, "field %v must be greater than or equal to %v", p.Field, minNumber)
			}
		} else if p.RValue.CanFloat() {
			valueFloat := p.RValue.Float()
			if valueFloat < minNumber {
				return NewValidationError(p.Field, "field %v must be greater than or equal to %v", p.Field, minNumber)
			}
		}
		return NilValidationError()
	}
}

func Max(maxNumber float64) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if p.RValue.CanInt() {
			valueInt := p.RValue.Int()
			if float64(valueInt) > maxNumber {
				return NewValidationError(p.Field, "field %v must be less than or equal to %v", p.Field, maxNumber)
			}
		} else if p.RValue.CanFloat() {
			valueFloat := p.RValue.Float()
			if valueFloat > maxNumber {
				return NewValidationError(p.Field, "field %v must be less than or equal to %v", p.Field, maxNumber)
			}
		}
		return NilValidationError()
	}
}
