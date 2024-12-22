package goi

import (
	"fmt"
	"reflect"
	"strconv"
)

func DecodeCustom(data map[string]any, v any) error {
	// Obtener el valor de la estructura a partir de la interfaz vacía
	val := reflect.ValueOf(v).Elem()
	typ := reflect.TypeOf(v).Elem()

	// Recorrer los campos de la struct
	for i := 0; i < val.NumField(); i++ {
		// Obtener el campo actual y su etiqueta "x"
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("goi") // Obtener la etiqueta personalizada "x"

		// Verificar si la etiqueta personalizada existe en los datos
		if value, exists := data[tag]; exists {
			// Verificar el tipo del campo y asignar el valor
			switch field.Kind() {
			case reflect.String:
				if strVal, ok := value.(string); ok {
					field.SetString(strVal)
				} else {
					return fmt.Errorf("expected string for %s, got %T", tag, value)
				}
			case reflect.Int:
				if intVal, ok := value.(int); ok {
					field.SetInt(int64(intVal))
				} else if strVal, ok := value.(string); ok {
					// Si es una cadena de texto, intentamos convertirla a int
					parsedInt, err := strconv.Atoi(strVal)
					if err != nil {
						return fmt.Errorf("failed to convert string '%s' to int", strVal)
					}
					field.SetInt(int64(parsedInt))
				} else {
					return fmt.Errorf("expected int for %s, got %T", tag, value)
				}
			case reflect.Bool:
				if boolVal, ok := value.(bool); ok {
					field.SetBool(boolVal)
				} else if strVal, ok := value.(string); ok {
					// Intentamos convertir de string a bool
					parsedBool, err := strconv.ParseBool(strVal)
					if err != nil {
						return fmt.Errorf("failed to convert string '%s' to bool", strVal)
					}
					field.SetBool(parsedBool)
				} else {
					return fmt.Errorf("expected bool for %s, got %T", tag, value)
				}
			default:
				return fmt.Errorf("unsupported type %s for field %s", field.Kind(), fieldType.Name)
			}
		}
	}

	return nil
}