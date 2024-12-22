package goi

import (
	"fmt"
	"reflect"
)

type NumberValidator struct {
	Validator
}

func Number() *NumberValidator {
	v := &NumberValidator{}
	v.NumberBase()
	return v
}

func (s *NumberValidator) NumberBase() {
	s.ruleNames = append(s.ruleNames, "base")
	s.rules = append(s.rules, func(value *any) error {
		if reflect.TypeOf(*value).Kind() != reflect.Float64 && reflect.TypeOf(*value).Kind() != reflect.Float32 {
			return fmt.Errorf("%s not a number", s.name)
		}
		return nil
	})
}

func (s *NumberValidator) Required() *NumberValidator {
	s.ruleNames = append(s.ruleNames, "required")
	// s.rules = append(s.rules, func(value *any) error {
	// 	if value == nil {
	// 		return fmt.Errorf("%s required2", s.name)
	// 	}
	// 	return nil
	// })
	return s
}

func (s *NumberValidator) Min(v float64) *NumberValidator {
	s.ruleNames = append(s.ruleNames, "min")
	s.rules = append(s.rules, func(value *any) error {
		cast := (*value).(float64)
		if cast < v {
			return fmt.Errorf("%s must be greater than or equal to %.0f", s.name, v)
		}
		return nil
	})
	return s
}

func (s *NumberValidator) Validate(data *any) error {
	val := reflect.ValueOf(*data)
	if val.Kind() != reflect.Map {
		return fmt.Errorf("not a map")
	}
	for _, k := range val.MapKeys() {
		s.name = k.String()
		v := val.MapIndex(k).Interface()
		//TODO: Validate optional validation
		if v == nil && findIndex(s.ruleNames, "required") != -1 {
			return fmt.Errorf("%s required", s.name)
		}
		s.value = &v
		for _, rule := range s.rules {
			if err := rule(&v); err != nil {
				return err
			}
		}
	}
	return nil
}
