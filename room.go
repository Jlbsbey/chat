package main

import (
	"encoding/json"
	"fmt"
)

type Room struct {
	Name     string            `json:"name"`
	Messages []Message         `json:"messages"`
	ID       string            `json:"id"`
	Users    map[uint64]uint32 `json:"allowedUsersID"`
}

type CreateRoom struct {
	R Room `json:"data"`
}
type EditRoom struct {
	R Room `json:"data"`
}
type ReadRoom struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}
type DeleteRoom struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (r Room) Create() DefinedAction {
	return &CreateRoom{}
}
func (action *CreateRoom) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action CreateRoom) Process(db *DB) {
	action.R.ID = fmt.Sprint(FirstFreeID)
	FirstFreeID++
	*db = append(*db, action.R)
}

func (r Room) Edit() DefinedAction {
	return &EditRoom{}
}
func (action *EditRoom) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action EditRoom) Process(db *DB) {
	id := action.R.GetID()
	(*db)[db.GetIndex(id)] = action.R
}

func (r Room) Delete() DefinedAction {
	return &DeleteRoom{}
}
func (action *DeleteRoom) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action DeleteRoom) Process(db *DB) {
	for i, p := range *db {
		if p.GetID() == action.Data.ID {
			*db = append((*db)[:i], (*db)[i+1:]...)
		}
	}
}

func (r Room) Read() DefinedAction {
	return &ReadRoom{}
}
func (action *ReadRoom) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadRoom) Process(db *DB) {
	(*db)[db.GetIndex(action.Data.ID)].Print()
}

func (r Room) Print() {
	fmt.Println("Name", r.Name, ", ID", r.ID, ", Messages", r.Messages, ", Users", r.Users)
}
func (r Room) GetID() string {
	return r.ID
}
