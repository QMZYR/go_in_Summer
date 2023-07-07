package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Username string
	Password string
}

var users = make(map[string]User)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("请选择操作：")
		fmt.Println("1. 注册")
		fmt.Println("2. 登录")
		fmt.Println("3. 修改密码")
		fmt.Println("4. 退出")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			register(scanner)
		case "2":
			login(scanner)
		case "3":
			changePassword(scanner)
		case "4":
			fmt.Println("程序已退出！")
			return
		default:
			fmt.Println("无效的选择，请重新输入！")
		}
	}
}

func register(scanner *bufio.Scanner) {
	fmt.Println("请输入用户名：")
	scanner.Scan()
	username := scanner.Text()

	if _, ok := users[username]; ok {
		fmt.Println("用户名已存在！")
		return
	}

	fmt.Println("请输入密码：")
	scanner.Scan()
	password := scanner.Text()

	users[username] = User{
		Username: username,
		Password: password,
	}

	fmt.Println("注册成功！")
}

func login(scanner *bufio.Scanner) {
	fmt.Println("请输入用户名：")
	scanner.Scan()
	username := scanner.Text()

	fmt.Println("请输入密码：")
	scanner.Scan()
	password := scanner.Text()

	user, ok := users[username]
	if !ok || user.Password != password {
		fmt.Println("用户名或密码错误！")
		return
	}

	fmt.Println("登录成功！")
}

func changePassword(scanner *bufio.Scanner) {
	fmt.Println("请输入用户名：")
	scanner.Scan()
	username := scanner.Text()

	user, ok := users[username]
	if !ok {
		fmt.Println("用户不存在！")
		return
	}

	fmt.Println("请输入原密码：")
	scanner.Scan()
	oldPassword := scanner.Text()

	if user.Password != oldPassword {
		fmt.Println("原密码错误！")
		return
	}

	fmt.Println("请输入新密码：")
	scanner.Scan()
	newPassword := scanner.Text()

	user.Password = newPassword
	users[username] = user

	fmt.Println("密码修改成功！")
}
