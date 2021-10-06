package util

import "golang.org/x/crypto/bcrypt"

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

// SetPassword 设置密码
func SetPassword(password string) (passwordDigest string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return password, err
	}
	PasswordDigest := string(bytes)
	return PasswordDigest, nil
}

// CheckPassword 校验密码
func CheckPassword(PasswordDigest, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(PasswordDigest), []byte(password))
	return err == nil
}
