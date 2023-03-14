package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	cfg()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	go sendToConnection(conn)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if string(buf[:n]) == "/::login" {
			n, _ = conn.Read(buf)
			RecAction(buf[:n])
			if logined == false {
				conn.Write([]byte("|::failed"))
			}
		} else if string(buf[:n]) == "/::create" {

		} else {
			fmt.Println(string(buf[:n]))
		}
		if err != nil {
			panic(err)
		}

	}
}

func sendToConnection(conn net.Conn) {
	var s string
	for {
		fmt.Scan(&s)
		_, err := conn.Write([]byte(s))
		if err != nil {
			panic(err)
		}
	}
}

func RecAction(text []byte) {
	logined = false
	var action Action
	err := json.Unmarshal(text, &action)
	if err != nil {
		fmt.Println(err)
		return
	}
	var obj GeneralObject
	switch action.ObjName {
	case "user":
		obj = &User{}
	/*case "room":
	obj = &Room{}*/
	default:
		fmt.Println("Unknown object", action.ObjName)
	}

	var defact DefinedAction

	switch action.Action {
	case "create":
		obj.Create()
	case "edit":
		obj.Edit()
	case "delete":
		obj.Delete()
	case "read":
		obj.Read()
	case "login":
		defact = obj.Login()
	default:
		fmt.Println("Unknown action", action.Action)
		return
	}
	defact.GetFromJSON(text)
	defact.Process()
}
