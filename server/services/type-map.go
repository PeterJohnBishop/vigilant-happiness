package services

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func GenerateGoTypeMap(structName string, payload map[string]interface{}) (string, error) {
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

		sb.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", fieldName, fieldType, key))

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
		f := reflect.ValueOf(val).Float()
		if f == float64(int(f)) {
			return "int", ""
		}
		return "float64", ""
	case []interface{}:
		if len(val) == 0 {
			return "[]any", ""
		}

		first := val[0]
		elemType, nested := detectType(parentName, first, structs)

		if _, ok := first.(map[string]interface{}); ok {
			elemStructName := singularize(parentName)
			elemBody := buildNestedStruct(first.(map[string]interface{}), structs)
			structs[elemStructName] = elemBody
			return "[]" + elemStructName, ""
		}

		if nested != "" {
			structs[elemType] = nested
		}
		return "[]" + elemType, ""
	case map[string]interface{}:
		structName := parentName
		body := buildNestedStruct(val, structs)
		return structName, body
	default:
		return "any", ""
	}
}

func buildNestedStruct(payload map[string]interface{}, structs map[string]string) string {
	var sb strings.Builder
	for k, v2 := range payload {
		fieldName := toCamelCase(k)
		fieldType, nested := detectType(fieldName, v2, structs)
		sb.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", fieldName, fieldType, k))
		if nested != "" {
			structs[fieldType] = nested
		}
	}
	return sb.String()
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

// converts plural names to singular for type
func singularize(name string) string {
	if strings.HasSuffix(name, "s") {
		return name[:len(name)-1]
	}
	return name
}
