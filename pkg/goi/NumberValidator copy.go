package goi

// import (
// 	"fmt"
// 	"reflect"
// )

// type NumberValidator struct {
// 	Validator
// }

// func Number() *NumberValidator {
// 	v := &NumberValidator{}
// 	v.NumberBase()
// 	return v
// }

// func (s *NumberValidator) NumberBase() {
// 	s.ruleNames = append(s.ruleNames, "base")
// 	s.rules = append(s.rules, func(value *any) error {
// 		if reflect.TypeOf(*value).Kind() != reflect.Float64 && reflect.TypeOf(*value).Kind() != reflect.Float32 {
// 			return fmt.Errorf("%s not a number", s.label)
// 		}
// 		return nil
// 	})
// }

// func (s *NumberValidator) Required() *NumberValidator {
// 	s.ruleNames = append(s.ruleNames, "required")
// 	s.rules = append(s.rules, func(value *any) error {
// 		if value == nil || *value == nil {
// 			return fmt.Errorf("%s must be defined", s.label)
// 		}
// 		return nil
// 	})
// 	return s
// }

// func (s *NumberValidator) Min(v float64) *NumberValidator {
// 	s.ruleNames = append(s.ruleNames, "min")
// 	s.rules = append(s.rules, func(value *any) error {
// 		cast := (*value).(float64)
// 		if cast < v {
// 			return fmt.Errorf("%s must be greater than or equal to %.0f", s.label, v)
// 		}
// 		return nil
// 	})
// 	return s
// }

// func (s *NumberValidator) Default(defaultValue any) *NumberValidator {
// 	s.ruleNames = append(s.ruleNames, "default")
// 	s.defaultValue = defaultValue
// 	return s
// }

// func (s *NumberValidator) Optional() *NumberValidator {
// 	s.ruleNames = append(s.ruleNames, "optional")
// 	return s
// }

// func (s *NumberValidator) Label(label string) *NumberValidator {
// 	s.label = label
// 	return s
// }

// func (s *NumberValidator) Validate(data *any) error {
// 	if s.label == "" {
// 		s.label = "value"
// 	}
// 	for _, rule := range s.rules {
// 		if err := rule(data); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
