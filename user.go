package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type UserTemp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	Email    string `json:"email"`
}

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
func (action LoginUser) Process(db *DB) {
	folderPath := "C:\\program1\\go_spec\\chat\\users"
	fmt.Println("login2")
	count := 0
	filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".json") {
			count++
		}
		return nil
	})
	var i int
	for i = 1; i <= count; i++ {
		jsonFile, err := ioutil.ReadFile("users/user" + strconv.Itoa(i) + ".json")
		if err != nil {
			fmt.Println(err)
			return
		}
		var json_user User
		err = json.Unmarshal(jsonFile, &json_user)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(json_user.Username)
		fmt.Println(action.Data.Username)
		fmt.Println(json_user.Password)
		fmt.Println(action.Data.Password)
		if json_user.Username == action.Data.Username && json_user.Password == action.Data.Password {
			fmt.Println("Login successful")
			return
		} else {
			fmt.Println("Login failed")
			return
		}

	}
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
func (action CreateUser) Process(db *DB) {
	action.U.ID = fmt.Sprint(FirstFreeID)
	FirstFreeID++
	*db = append(*db, action.U)
}
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
func (action EditUser) Process(db *DB) {
	id := action.U.GetID()
	(*db)[db.GetIndex(id)] = action.U
}

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
func (action DeleteUser) Process(db *DB) {
	for i, p := range *db {
		if p.GetID() == action.Data.ID {
			*db = append((*db)[:i], (*db)[i+1:]...)
		}
	}
}

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
func (action ReadUser) Process(db *DB) {
	(*db)[db.GetIndex(action.Data.ID)].Print()
}

func (u User) Print() {
	fmt.Printf("Name: %s, Password: %s, ID: %d, Email: %s, Room: %v", u.Username, u.Password, u.ID, u.Email)
}
func (u User) GetID() string {
	return u.ID
}
