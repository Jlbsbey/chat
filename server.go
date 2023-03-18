package main

import (
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

func response(obj string, act string, state string, conn net.Conn) {
	fmt.Println("|::" + obj + "_::_" + act + "_::_" + state)
	conn.Write([]byte("|::" + obj + "_::_" + act + "_::_" + state))
}
