package autosdk

import (
	"strings"
	"unicode"
)

/*
{
	"refunds_receive_get_response": {
	    "refunds": {
	        "refund": [
	            {
	                "refund_id": 38046356523485
	            },
	            {
	                "refund_id": 36094501861333
	            }
	        ]
	    }
	}
}

*/
func ResponeField(name, _type, level string) string {
	n := GetHump(name)
	t := parseType(_type, level)
	var s = n + " " + t + `'json:"` + name + Float64toString(_type) + `"'`
	if level == "Object Array" {
		s = n + ` struct{` +
			_type + " " + t + ` 'json:"` + UnderlineTheHump(_type) + Float64toString(_type) + `"'` +
			`}'json:"` + n + `"'`
		// Refunds struct{
		// 	Refund []*Refund`json:"refund"`
		// }`json:"refunds"`
	}
	// fmt.Println(strings.Replace(s, "'", "`", -1))
	return strings.Replace(s, "'", "`", -1)
}
func Float64toString(_type string) string {
	if isFloat64(_type) {
		return ",string"
	}
	return ""
}

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
func xmlName(langType string) string {
	prefix := "[]*"
	if strings.HasPrefix(langType, prefix) {
		xmlName := strings.TrimPrefix(langType, prefix)
		runes := UnderlineTheHump(xmlName)
		return ">" + string(runes)
	}
	return ""
}

/*
驼峰转加下划线
*/
func UnderlineTheHump(s string) string {
	runes := []rune{}
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i != 0 {
				runes = append(runes, '_')
			}
			runes = append(runes, unicode.ToLower(r))
		} else {
			runes = append(runes, r)
		}
	}
	return string(runes)
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
