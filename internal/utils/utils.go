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

func ansi256ToRGB(n int) (r, g, b int) {
	if n < 16 {
		return 0, 0, 0
	}
	if n >= 232 {
		gray := (n-232)*10 + 8
		return gray, gray, gray
	}

	n -= 16
	r = (n / 36) * 51
	g = ((n % 36) / 6) * 51
	b = (n % 6) * 51
	return
}

func IsLightANSI(n int) bool {
	r, g, b := ansi256ToRGB(n)

	luminance :=
		0.299*float64(r) +
			0.587*float64(g) +
			0.114*float64(b)

	return luminance > 186
}
