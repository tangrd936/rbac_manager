package test_file

import (
	"fmt"
	"rbac_manager/utils/pwd"
	"testing"
)

// 测试密码加密
func TestHashPassword(t *testing.T) {
	hashPassword := pwd.HashPassword("123456")
	fmt.Println(hashPassword)
	ok := pwd.ComparePasswords(hashPassword, "123456")
	fmt.Println(ok)
}
