package utils

import "fmt"

func AnyAsString(v any) string {
	if v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	default:
		return fmt.Sprintf("%v", t)
	}
}

func AnyAsIntPtr(v any) *int {
	if v == nil {
		return nil
	}
	if f, ok := v.(float64); ok {
		n := int(f)
		return &n
	}
	if n, ok := v.(int); ok {
		return &n
	}
	return nil
}
