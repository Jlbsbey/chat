package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string          `json:"name"`
	Password string          `json:"password"`
	ID       string          `json:"id"`
	Email    string          `json:"email"`
	Room     map[uint64]bool `json:"room_id"`
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
	fmt.Printf("Name: %s, Password: %s, ID: %d, Email: %s, Room: %v", u.Name, u.Password, u.ID, u.Email, u.Room)
}
func (u User) GetID() string {
	return u.ID
}
