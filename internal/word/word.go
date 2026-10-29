package word

import (
	"strings"
	"unicode"
)

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func UnderscoreToUpperCamel(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}

func UnderscoreToLowerCamel(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if i == 0 {
			parts[i] = strings.ToLower(part)
		} else {
			parts[i] = strings.Title(part)
		}
	}
	return strings.Join(parts, "")
}

func CamelCaseToUnderscore(s string) string {
	// govalidator 库
	// just_a_var
	// 关于append，为什么不是 result.append()？因为 slice本质结构体，是纯值传递，所以不能直接进行操作。
	var result []rune
	for i, r := range s {
		if i == 0 {
			result = append(result, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}
