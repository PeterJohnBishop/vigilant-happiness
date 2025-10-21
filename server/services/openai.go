package services

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func GenerateTypeMap(payload map[string]interface{}) (string, error) {
	typeMap := make(map[string]string)

	for key, value := range payload {
		typeMap[key] = detectType(value)
	}

	result, err := json.MarshalIndent(typeMap, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling result: %w", err)
	}

	return string(result), nil
}

func detectType(v interface{}) string {
	switch val := v.(type) {
	case string:
		return "string"
	case bool:
		return "bool"
	case int, int8, int16, int32, int64:
		return "int"
	case float32, float64:
		if float64(int(val.(float64))) == val.(float64) {
			return "int"
		}
		return "float64"
	case []interface{}:
		if len(val) == 0 {
			return "[]any"
		}
		elemType := detectType(val[0])
		return "[]" + elemType
	case map[string]interface{}:
		nested := make(map[string]string)
		for k, v2 := range val {
			nested[k] = detectType(v2)
		}
		nestedJSON, _ := json.MarshalIndent(nested, "", "  ")
		return string(nestedJSON)
	default:
		return fmt.Sprintf("%s", reflect.TypeOf(v))
	}
}
