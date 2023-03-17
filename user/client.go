package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       string `json:"id"`
	Email    string `json:"email"`
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	handleConnections(conn)
}

func handleConnections(conn net.Conn) {
	go sendToConnection(conn)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		switch string(buf[:n]) {
		case "|::user_::_login_::_fail":
			fmt.Println("Login failed")
			return
		case "|::user_::_login_::_succ":
			fmt.Println("Login successful")
		default:
			fmt.Println(string(buf[:n]))
		}
	}
}

func sendToConnection(conn net.Conn) {
	var s string
	fmt.Print("Choose an action: /login or /create: ")
	fmt.Scan(&s)
	if s == "/login" {
		login(conn)
	} else if s == "/create" {
		create(conn)
	} else {
		fmt.Println("Unknown action")
		return
	}
	for {
		fmt.Scan(&s)
		_, err := conn.Write([]byte(s))
		if err != nil {
			panic(err)
		}
	}
}

func login(conn net.Conn) {
	conn.Write([]byte("/::login"))
	var username string
	var password string
	fmt.Print("Write a login: ")
	fmt.Scan(&username)
	fmt.Print("Write a password: ")
	fmt.Scan(&password)
	u, err := json.Marshal(Action{Action: "login", ObjName: "user", Data: User{Username: username, Password: password}})
	if err != nil {
		panic(err)
	}
	conn.Write(u)
}

func create(conn net.Conn) {
	var username, password, email string
	conn.Write([]byte("/::create"))
	fmt.Print("Write a login: ")
	fmt.Scan(&username)
	fmt.Print("Write a password: ")
	fmt.Scan(&password)
	fmt.Print("Write a email(optional, leave empty if do not want to write it): ")
	fmt.Scan(&email)
	u, err := json.Marshal(Action{Action: "create", ObjName: "user", Data: User{Username: username, Password: password, Email: email}})
	if err != nil {
		panic(err)
	}
	conn.Write(u)
}
