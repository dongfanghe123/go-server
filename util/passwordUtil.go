package util

import "golang.org/x/crypto/bcrypt"

// 1. 生成密码哈希 (注册时)
func HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost 是 10，这是一个性能和安全性的良好平衡点[citation:6]
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// 2. 验证密码 (登录时)
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // 如果密码匹配，err 为 nil[citation:2]
}
