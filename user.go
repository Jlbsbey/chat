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
	ID       int    `json:"user_id"`
	Email    string `json:"email"`
}

type CreateUser struct {
	U User `json:"data"`
}
type EditUser struct {
	Data struct {
		Session_ID uint64 `json:"session_id"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		User_ID    int    `json:"user_id"`
		Email      string `json:"email"`
	}
}
type ReadUser struct {
	Data struct {
		Session_ID uint64 `json:"session_id"`
	} `json:"data"`
}
type DeleteUser struct {
	Data struct {
		Password   string `json:"password"`
		Session_ID uint64 `json:"session_id"`
		User_ID    int    `json:"user_id"`
	} `json:"data"`
}

type LogoutUser struct {
	Data struct {
		Session_ID uint64 `json:"session_id"`
	}
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
	var login, passw, email string
	var ID int
	for lg.Next() {
		if err = lg.Scan(&login, &passw); err != nil {
			log.Println(err)
		}
		if login == action.Data.Username && passw == action.Data.Password {
			logined = true
			us = `SELECT ID, Login, Email FROM users WHERE Login = ?`
			lg, err = db.Query(us, action.Data.Username)
			for lg.Next() {
				if err = lg.Scan(&ID, &login, &email); err != nil {
					log.Println(err)
				}
			}
			if err != nil {
				panic(err)
			}
			ses_id := CreateSession(ID)
			response, err := json.Marshal(Response{Session_ID: ses_id, Action: "login", Success: true, ObjName: "user", Data: User{Username: login, ID: ID, Email: email}})
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
	var count int
	quer := `SELECT Login FROM users WHERE Login = ?`
	lg, err := db.Query(quer, action.U.Username)
	for lg.Next() {
		count++
	}
	if count > 0 {
		response, err := json.Marshal(Response{Action: "register", Success: false, ObjName: "user", Error_message: "User already exists"})
		if err != nil {
			panic(err)
		}
		return response
	}
	quer = `INSERT INTO users(Login, Password, Email) VALUES (?, ?, ?)`
	_, err = db.ExecContext(context.Background(), quer, action.U.Username, action.U.Password, action.U.Email)
	if err != nil {
		response, err := json.Marshal(Response{Action: "register", Success: false, ObjName: "user"})
		if err != nil {
			panic(err)
		}
		return response
	} else {
		response, err := json.Marshal(Response{Action: "register", Success: true, ObjName: "user"})
		if err != nil {
			panic(err)
		}

		return response
	}
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
func (action EditUser) Process() []byte {
	if sessions[action.Data.Session_ID] == action.Data.User_ID {
		us := `UPDATE users SET Login = ?, Password = ?, Email = ? WHERE ID = ?`
		_, err := db.ExecContext(context.Background(), us, action.Data.Username, action.Data.Password, action.Data.Email, sessions[action.Data.Session_ID])
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(Response{Action: "edit", Success: true, ObjName: "user"})
		return response
	} else {
		response, err := json.Marshal(Response{Action: "edit", Success: false, ObjName: "user", Error_message: "Access denied"})
		if err != nil {
			panic(err)
		}
		return response
	}

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
func (action DeleteUser) Process() []byte {
	us := `SELECT Password FROM users WHERE ID = ?`
	lg, err := db.Query(us, sessions[action.Data.Session_ID])
	if err != nil {
		panic(err)
	}
	var passw string
	for lg.Next() {
		if err = lg.Scan(&passw); err != nil {
			log.Println(err)
		}
		if passw == action.Data.Password {
			if sessions[action.Data.Session_ID] == action.Data.User_ID {
				us = `DELETE FROM users WHERE ID = ?`
				_, err := db.ExecContext(context.Background(), us, sessions[action.Data.Session_ID])
				if err != nil {
					panic(err)
				}
				us = `DELETE FROM users_rooms WHERE User_id = ?`
				_, err = db.ExecContext(context.Background(), us, sessions[action.Data.Session_ID])
				if err != nil {
					panic(err)
				}
				us = `UPDATE messages SET Author_ID = ? WHERE Author_ID = ?`
				_, err = db.ExecContext(context.Background(), us, -1, sessions[action.Data.Session_ID])
				if err != nil {
					panic(err)
				}
				response, err := json.Marshal(Response{Action: "delete", Success: true, ObjName: "user"})
				return response
			} else {
				response, err := json.Marshal(Response{Action: "delete", Success: false, ObjName: "user", Error_message: "Access denied"})
				if err != nil {
					panic(err)
				}
				return response
			}
		} else {
			response, err := json.Marshal(Response{Action: "delete", Success: false, ObjName: "user", Error_message: "Access denied"})
			if err != nil {
				panic(err)
			}
			return response
		}
	}
	response, err := json.Marshal(Response{Action: "delete", Success: false, ObjName: "user", Error_message: "Access denied"})
	if err != nil {
		panic(err)
	}
	return response
}

func (u User) Logout() DefinedAction {
	return &LogoutUser{}
}
func (action *LogoutUser) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action LogoutUser) Process() []byte { return nil }

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
func (action ReadUser) Process() []byte {
	var login, email string
	userID := sessions[action.Data.Session_ID]
	fmt.Println(userID)
	readQuery := `SELECT ID, Login, Email FROM users WHERE ID = ?`
	rows, err := db.Query(readQuery, userID)
	for rows.Next() {
		if err = rows.Scan(&userID, &login, &email); err != nil {
			log.Println(err)
		}
	}
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(Response{Action: "read", ObjName: "user", Data: User{Username: login, ID: userID, Email: email}})
	if err != nil {
		panic(err)
	}
	fmt.Println(5)
	return response

}

func (u User) Print() {
	fmt.Printf("Name: %s, Password: %s, ID: %d, Email: %s, Room: %v", u.Username, u.Password, u.ID, u.Email)
}
func (u User) GetID() int {
	return u.ID
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
	return id
}
