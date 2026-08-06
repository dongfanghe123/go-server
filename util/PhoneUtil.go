package util

import (
	"regexp"
)

// IsValidPhone 校验中国大陆手机号
func IsValidPhone(phone string) bool {
	// 规则：1开头，第二位3-9，后面9位数字
	reg := `^1[3-9]\d{9}$`
	matched, _ := regexp.MatchString(reg, phone)
	return matched
}
