package services

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func GenerateTypeMap(structName string, payload map[string]interface{}) (string, error) {
	structs := make(map[string]string)
	buildStruct(structName, payload, structs)

	var result strings.Builder
	for name, body := range structs {
		result.WriteString(fmt.Sprintf("type %s struct {\n%s}\n\n", name, body))
	}

	return result.String(), nil
}

func buildStruct(name string, payload map[string]interface{}, structs map[string]string) {
	var sb strings.Builder

	for key, value := range payload {
		fieldName := toCamelCase(key)
		fieldType, nestedStruct := detectType(fieldName, value, structs)

		sb.WriteString(fmt.Sprintf("    %s %s `json:\"%s\" `\n", fieldName, fieldType, key))

		if nestedStruct != "" {
			structs[fieldType] = nestedStruct
		}
	}

	structs[name] = sb.String()
}

func detectType(parentName string, v interface{}, structs map[string]string) (string, string) {
	switch val := v.(type) {
	case string:
		return "string", ""
	case bool:
		return "bool", ""
	case int, int8, int16, int32, int64:
		return "int", ""
	case float32, float64:
		if reflect.ValueOf(val).Float() == float64(int(reflect.ValueOf(val).Float())) {
			return "int", ""
		}
		return "float64", ""
	case []interface{}:
		if len(val) == 0 {
			return "[]any", ""
		}
		elemType, nested := detectType(parentName, val[0], structs)
		return "[]" + elemType, nested
	case map[string]interface{}:
		structName := parentName
		var sb strings.Builder
		for k, v2 := range val {
			fieldName := toCamelCase(k)
			fieldType, nested := detectType(fieldName, v2, structs)
			sb.WriteString(fmt.Sprintf("    %s %s `json:\"%s\" dynamodbav:\"%s\"`\n", fieldName, fieldType, k, k))
			if nested != "" {
				structs[fieldType] = nested
			}
		}
		return structName, sb.String()
	default:
		return "any", ""
	}
}

func toCamelCase(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})
	for i := range parts {
		runes := []rune(parts[i])
		runes[0] = unicode.ToUpper(runes[0])
		parts[i] = string(runes)
	}
	return strings.Join(parts, "")
}
