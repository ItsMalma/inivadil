package inivadil

func If(ruleElement RuleElement) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if err := ruleElement(p); !err.Nil() {
			return ValidationError{stop: true}
		}
		return NilValidationError()
	}
}

func IfNot(ruleElement RuleElement) RuleElement {
	return func(p RuleElementParameter) ValidationError {
		if err := ruleElement(p); err.Nil() {
			return ValidationError{stop: true}
		}
		return NilValidationError()
	}
}
