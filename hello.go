package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Account struct {
	Name     string
	Password string
}

// 结果，注意首字母大写，不然json不能解析
type Result struct {
	OK bool
}

func main() {
	// 当前注册的所有用户
	accounts := make([]*Account, 0, 6)
	// register下的处理
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		// 获取name和password
		name := r.FormValue("Name")
		pwd := r.FormValue("Password")

		// 判断内存内是否已经保存了该用户
		if HasAccountByName(accounts, name) {
			ReturnResult(false, w)
		} else {
			// 没有该用户则存入
			accounts = append(accounts, &Account{
				Name:     name,
				Password: pwd,
			})
			ReturnResult(true, w)
		}
	})
	// login下的处理
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// 获取name和password
		name := r.FormValue("Name")
		pwd := r.FormValue("Password")

		// 判断用户名密码是否匹配
		if IsAccountPwdMatch(accounts, name, pwd) {
			ReturnResult(true, w)
		} else {
			ReturnResult(false, w)
		}
	})

	// 开启服务器
	log.Println("Server start!")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

// 给客户端返回结果
func ReturnResult(result bool, w http.ResponseWriter) {
	// 创建结果类
	re := &Result{
		OK: result,
	}

	// 使用json构造字符串
	j, _ := json.Marshal(re)

	// 发送给客户端
	w.Write([]byte(j))
}

// 根据名称判断是否有该用户
func HasAccountByName(as []*Account, name string) bool {
	for _, v := range as {
		if v.Name == name {
			return true
		}
	}
	return false
}

// 根据名称和密码判断该用户是否匹配
func IsAccountPwdMatch(as []*Account, name string, password string) bool {
	for _, v := range as {
		if v.Name == name && v.Password == password {
			return true
		}
	}
	return false
}
