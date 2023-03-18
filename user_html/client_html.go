package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	var body bytes.Buffer
	if login(body, client) == false {
		return
	}

}

func login(body bytes.Buffer, client *http.Client) bool {
	var username, password string
	fmt.Print("Write a login: ")
	fmt.Scan(&username)
	fmt.Print("Write a password: ")
	fmt.Scan(&password)
	u, err := json.Marshal(Action{Action: "login", ObjName: "user", Data: User{Username: username, Password: password}})
	if err != nil {
		panic(err)
	}

	body.Write(u)
	req, err := http.NewRequest("POST", "http://localhost:8080/", &body)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	var response Response
	err = json.Unmarshal(data, &response)
	if response.Success == true {
		fmt.Println("Login successful")
		return true
	} else if response.Success == false {
		fmt.Println("Login failed")
		return false
	}
	return false
}
