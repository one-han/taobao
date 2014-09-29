package autosdk

import (
	"strings"
	"unicode"
)

/**
* 去掉字符中的".""_",然后转为驼峰
* 比如：user.location_name
* return UserLocationName
 */
func GetHump(s string) string {
	s = strings.Replace(s, ".", " ", -1)
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	return s
}

func GetRequestResponseStructName(s, suffix string) string {
	s = strings.TrimSuffix(s, suffix)
	s = strings.TrimPrefix(s, "taobao.")
	return GetHump(s) + suffix
}

func GetRequestResponseStructJsonName(s string) string {
	return strings.Replace(strings.TrimPrefix(s, "taobao."), ".", "_", -1) + "_response"
}
func hasObject(level string) bool {
	return strings.HasPrefix(level, "Object")
}
func isFloat64(_type string) bool {
	return _type == "Price"
}
func jsonName(langType string) string {
	prefix := "[]*"
	log.Println("langType：", langType)
	if strings.HasPrefix(langType, prefix) {
		log.Println("哈哈哈哈h")
		jsonName := strings.TrimPrefix(langType, prefix)
		runes := []rune{}
		for i, r := range jsonName {
			if unicode.IsUpper(r) {
				if i != 0 {
					runes = append(runes, '_')
				}
				runes = append(runes, unicode.ToLower(r))
			} else {
				runes = append(runes, r)
			}
		}
		return ">" + string(runes)
	}
	return ""
}
func parseType(_type, level string) (langType string) {
	switch _type {
	case "Number":
		langType = "int64"
	case "Price":
		langType = "float64"
	case "String", "Date", "Field List":
		langType = "string"
	case "Boolean":
		langType = "bool"
	case "byte[]":
		langType = "[]byte"
	default:
		langType = _type
	}

	switch level {
	case "Object":
		langType = "*" + langType
	case "Object Array":
		langType = "[]*" + langType
	case "Basic Array":
		langType = "[]" + langType
	}
	return
}
