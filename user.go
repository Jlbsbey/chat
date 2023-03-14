package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       string `json:"id"`
	Email    string `json:"email"`
}

type CreateUser struct {
	U User `json:"data"`
}
type EditUser struct {
	U User `json:"data"`
}
type ReadUser struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}
type DeleteUser struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

type LoginUser struct {
	Data struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		ID       string `json:"id"`
	} `json:"data"`
}

func (u User) Login() DefinedAction {
	return &LoginUser{}
}
func (action *LoginUser) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action LoginUser) Process(dB *DB) {
	us := `SELECT Login, Password FROM users WHERE Login = ? AND Password = ?`

	lg, err := db.Query(us, action.Data.Username, action.Data.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	var login, passw string
	for lg.Next() {
		if err = lg.Scan(&login, &passw); err != nil {
			log.Println(err)
		}
		if login == action.Data.Username && passw == action.Data.Password {
			logined = true
			fmt.Println("Login successful")
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Login failed")
	return
}
func (u User) Create() DefinedAction {
	return &CreateUser{}
}
func (action *CreateUser) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action CreateUser) Process(dB *DB) {}

func (u User) Edit() DefinedAction {
	return &EditUser{}
}
func (action *EditUser) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action EditUser) Process(dB *DB) {}

func (u User) Delete() DefinedAction {
	return &DeleteUser{}
}
func (action *DeleteUser) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action DeleteUser) Process(dB *DB) {}

func (u User) Read() DefinedAction {
	return &ReadUser{}
}
func (action *ReadUser) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadUser) Process(dB *DB) {}

func (u User) Print() {
	fmt.Printf("Name: %s, Password: %s, ID: %d, Email: %s, Room: %v", u.Username, u.Password, u.ID, u.Email)
}
func (u User) GetID() string {
	return u.ID
}
