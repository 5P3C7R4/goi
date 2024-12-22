package goi

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

type StringValidator struct {
	Validator
}

func String() *StringValidator {
	v := &StringValidator{}
	v.StringBase()
	return v
}

func (s *StringValidator) StringBase() {
	s.ruleNames = append(s.ruleNames, "base")
	s.rules = append(s.rules, func(value *any) error {
		if reflect.TypeOf(*value).Kind() != reflect.String {
			return fmt.Errorf("%s not a string", s.name)
		}
		return nil
	})
}

func (s *StringValidator) Required() *StringValidator {
	s.ruleNames = append(s.ruleNames, "required")
	s.rules = append(s.rules, func(value *any) error {
		cast, _ := (*value).(string)
		if cast == "" {
			return fmt.Errorf("%s required", s.name)
		}
		return nil
	})
	return s
}

func (s *StringValidator) LowerCase() *StringValidator {
	s.ruleNames = append(s.ruleNames, "lowercase")
	s.rules = append(s.rules, func(value *any) error {
		val := (any)(strings.ToLower((*value).(string)))
		*s.value = val
		*value = val
		return nil
	})
	return s
}

func (s *StringValidator) Trim() *StringValidator {
	s.ruleNames = append(s.ruleNames, "trim")
	s.rules = append(s.rules, func(value *any) error {
		val := (any)(strings.TrimSpace((*value).(string)))
		*s.value = &val
		return nil
	})
	return s
}

func (s *StringValidator) Min(length int) *StringValidator {
	s.ruleNames = append(s.ruleNames, "min")
	s.rules = append(s.rules, func(value *any) error {
		cast, _ := (*value).(string)
		if utf8.RuneCountInString(cast) < length {
			return fmt.Errorf("%s must be at least %d length", s.name, length)
		}
		return nil
	})
	return s
}

func (s *StringValidator) Validate(data *any) error {
	val := reflect.ValueOf(*data)
	fmt.Println(reflect.Va)
	if val.Kind() != reflect.Map {
		return fmt.Errorf("not a map")
	}
	for _, k := range val.MapKeys() {
		s.name = k.String()
		v := val.MapIndex(k).Interface()
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
