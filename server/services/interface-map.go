package services

import (
	"fmt"
	"reflect"
	"strings"
)

func GenerateInterfaceMap(interfaceName string, payload map[string]interface{}) (string, error) {
	interfaces := make(map[string]string)
	buildInterface(interfaceName, payload, interfaces)

	var result strings.Builder
	for name, body := range interfaces {
		displayName := toCamelCase(name)
		result.WriteString(fmt.Sprintf("interface %s {\n%s}\n\n", displayName, body))
	}

	return result.String(), nil
}

func buildInterface(name string, payload map[string]interface{}, interfaces map[string]string) {
	var sb strings.Builder

	for key, value := range payload {
		fieldType, nestedInterface := detectInterfaceTypes(key, value, interfaces)

		sb.WriteString(fmt.Sprintf("    %s: %s;\n", key, fieldType))

		if nestedInterface != "" {
			interfaces[fieldType] = nestedInterface
		}
	}

	interfaces[name] = sb.String()
}

func detectInterfaceTypes(parentName string, v interface{}, interfaces map[string]string) (string, string) {
	switch val := v.(type) {
	case string:
		return "string", ""
	case bool:
		return "bool", ""
	case int, int8, int16, int32, int64:
		return "number", ""
	case float32, float64:
		f := reflect.ValueOf(val).Float()
		if f == float64(int(f)) {
			return "number", ""
		}
		return "number", ""
	case []interface{}:
		if len(val) == 0 {
			return "[]any", ""
		}

		first := val[0]
		elemType, nested := detectType(parentName, first, interfaces)

		if _, ok := first.(map[string]interface{}); ok {
			elemStructName := singularize(parentName)
			elemBody := buildExtendingInterface(first.(map[string]interface{}), interfaces)
			interfaces[elemStructName] = elemBody
			elemStructDisplayName := toCamelCase(elemStructName)
			return elemStructDisplayName + "[]", ""
		}

		if nested != "" {
			interfaces[elemType] = nested
		}
		elemTypeDisplayName := toCamelCase(elemType)
		return elemTypeDisplayName + "[]", ""
	case map[string]interface{}:
		structName := parentName
		body := buildExtendingInterface(val, interfaces)
		displayName := toCamelCase(structName)
		return displayName, body
	default:
		return "any", ""
	}
}

func buildExtendingInterface(payload map[string]interface{}, structs map[string]string) string {
	var sb strings.Builder
	for k, v2 := range payload {
		fieldType, nested := detectInterfaceTypes(k, v2, structs)
		sb.WriteString(fmt.Sprintf("    %s: %s;\n", k, fieldType))
		if nested != "" {
			structs[fieldType] = nested
		}
	}
	return sb.String()
}
