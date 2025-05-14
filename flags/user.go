package flags

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"rbac_manager/global"
	"rbac_manager/models"
	"rbac_manager/utils/pwd"
)

type User struct {
}

func (u *User) CreateAdminUser() {
	fmt.Println("please input username:")
	var userName string
	_, err := fmt.Scanln(&userName)
	if err != nil {
		global.Log.Error(fmt.Sprintf("input username error: %v", err.Error()))
		return
	}
	var user models.UserModel
	err = global.Db.Take(&user, "user_name=?", userName).Error
	if err == nil {
		global.Log.Error(fmt.Sprintf("username is exist"))
		return
	}

	fmt.Println("please input password:")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		global.Log.Error(fmt.Sprintf("input password error: %v", err.Error()))
		return
	}
	fmt.Println("please input password again:")
	passwordT, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		global.Log.Error(fmt.Sprintf("again input password error: %v", err.Error()))
		return
	}
	if string(password) != string(passwordT) {
		global.Log.Error("The passwords entered twice are inconsistent")
		return
	}

	hashPassword := pwd.HashPassword(string(password))
	err = global.Db.Create(&models.UserModel{
		UserName: userName,
		Password: hashPassword,
		IsAdmin:  true,
	}).Error
	if err != nil {
		global.Log.Error("Create admin user error: " + err.Error())
		return
	}
	global.Log.Error("Create admin user success")
}
