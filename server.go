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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	go sendToConnection(conn)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		switch string(buf[:n]) {
		case "/::login":
			n, _ = conn.Read(buf)
			RecAction(buf[:n])
			if logined == false {
				return
				response("user", "login", "fail", conn)
			} else {
				response("user", "login", "succ", conn)
			}
		case "/::create":
			n, _ = conn.Read(buf)
			RecAction(buf[:n])
		default:
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
		defact = obj.Create()
	case "edit":
		defact = obj.Edit()
	case "delete":
		defact = obj.Delete()
	case "read":
		defact = obj.Read()
	case "login":
		defact = obj.Login()
	default:
		fmt.Println("Unknown action", action.Action)
		return
	}
	defact.GetFromJSON(text)
	defact.Process()

}

func response(obj string, act string, state string, conn net.Conn) {
	conn.Write([]byte("|::" + obj + "_::_" + act + "_::_" + state))
}
