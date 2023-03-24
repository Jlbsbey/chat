package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
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
		Login    bool   `json:"login"`
	} `json:"data"`
}

var sessions = make(map[uint64]int)

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
func (action LoginUser) Process() []byte {
	us := `SELECT Login, Password FROM users WHERE Login = ? AND Password = ?`

	lg, err := db.Query(us, action.Data.Username, action.Data.Password)
	if err != nil {
		panic(err)
	}
	var login, passw string
	var ID int
	for lg.Next() {
		if err = lg.Scan(&login, &passw); err != nil {
			log.Println(err)
		}
		if login == action.Data.Username && passw == action.Data.Password {
			logined = true
			us = `SELECT ID, Login FROM users WHERE Login = ?`
			lg, err = db.Query(us, action.Data.Username)
			if err != nil {
				panic(err)
			}
			lg.Scan(&ID)
			ses_id := CreateSession(ID)
			response, err := json.Marshal(Response{Session_ID: ses_id, Action: "login", Success: true, ObjName: "user"})
			if err != nil {
				panic(err)
			}
			return response
		}
		if err != nil {
			panic(err)
		}
	}
	response, err := json.Marshal(Response{Action: "login", Success: false, ObjName: "user"})
	return response
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
func (action CreateUser) Process() []byte {
	quer := `INSERT INTO users(Login, Password, Email) VALUES (?, ?, ?)`
	_, err := db.ExecContext(context.Background(), quer, action.U.Username, action.U.Password, action.U.Email)
	if err != nil {
		response, err := json.Marshal(Response{Action: "create", Success: false, ObjName: "user"})
		if err != nil {
			panic(err)
		}
		return response
	} else {
		response, err := json.Marshal(Response{Action: "create", Success: true, ObjName: "user"})
		if err != nil {
			panic(err)
		}

		return response
	}
}

/*func (u User) Edit() DefinedAction {
	return &EditUser{}
}
func (action *EditUser) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action EditUser) Process() {}

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
func (action DeleteUser) Process() {}

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
func (action ReadUser) Process() {}

func (u User) Print() {
	fmt.Printf("Name: %s, Password: %s, ID: %d, Email: %s, Room: %v", u.Username, u.Password, u.ID, u.Email)
}
func (u User) GetID() int {
	return u.ID
}*/

func first_id() int {
	var next_id int
	query := `SELECT MIN(ID+1) AS next_id FROM users WHERE ID+1 NOT IN (SELECT ID FROM users)`
	err := db.QueryRow(query).Scan(&next_id)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return next_id
}

func CreateSession(us_id int) uint64 {
	var id uint64
	var free bool
	free = false
	var max, min int
	max = 1000000000000000
	min = 100000000000000
	for free == false {
		rand.Seed(time.Now().UnixNano())
		id = uint64(rand.Intn(max-min) + min)
		if _, ok := sessions[id]; !ok {
			free = true
		}
	}
	sessions[id] = us_id
	fmt.Println(us_id)
	fmt.Println(id)
	return id
}
