package goi

import (
	"fmt"
	"reflect"
)

type SchemaValidator struct {
	schemaRuleNamesIn []string
	schemasIn         []any

	schemaValueNames []string
	values           []any
}

func Schema(m map[string]any) *SchemaValidator {
	s := &SchemaValidator{}
	for k, v := range m {
		s.schemaRuleNamesIn = append(s.schemaRuleNamesIn, k)
		s.schemasIn = append(s.schemasIn, v)
	}
	return s
}

func (s *SchemaValidator) Validate(data *any) error {
	dataVal := reflect.ValueOf(*data)
	if dataVal.Kind() != reflect.Map {
		return fmt.Errorf("not a map")
	}

	for _, k := range dataVal.MapKeys() {
		s.schemaValueNames = append(s.schemaValueNames, k.String())
		v := dataVal.MapIndex(k).Interface()
		s.values = append(s.values, v)
	}

	for i, v := range s.schemasIn {
		fieldName := s.schemaRuleNamesIn[i]
		index := findIndex(s.schemaValueNames, fieldName)
		var1 := (any)(map[string]any{s.schemaValueNames[index]: s.values[index]})
		valMethod := reflect.ValueOf(v).MethodByName("Validate")
		if !valMethod.IsValid() {
			panic("Method validate not found")
		}
		result := valMethod.Call([]reflect.Value{reflect.ValueOf(&var1)})
		if len(result) > 0 && result[0].Interface() != nil {
			return result[0].Interface().(error)
		}

	}

	return nil
}
