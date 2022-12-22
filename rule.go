package inivadil

import "reflect"

type RuleElementParameter struct {
	Field  string
	Value  any
	RValue reflect.Value
}

type RuleElement func(p RuleElementParameter) ValidationError

type Rule []RuleElement

type Rules map[string]Rule
