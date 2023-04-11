package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Room struct {
	Name string `json:"name"`
	//Messages   []Message `json:"messages"`
	Room_ID    string `json:"room_id"`
	InviteCode int    `json:"inv_code"`
}

type LoginRoom struct {
	Data struct {
		Name       string `json:"name"`
		ID         string `json:"id"`
		InviteCode int    `json:"invite_code"`
		User_ID    int    `json:"user_id"`
	} `json:"data"`
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

func (r Room) Login() DefinedAction {
	return &LoginRoom{}
}
func (action *LoginRoom) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action LoginRoom) Process() []byte {
	us := `SELECT Room_name, Invite_code FROM rooms WHERE Room_name = ? OR Invite_code = ?`
	fmt.Println(action.Data.Name)
	fmt.Println(action.Data.InviteCode)
	fmt.Println(action.Data.User_ID)
	lg, err := db.Query(us, action.Data.Name, action.Data.InviteCode)
	fmt.Println(lg)
	if err != nil {
		panic(err)
	}
	var name string
	var inv_c int
	fmt.Println(3)
	for lg.Next() {
		if err = lg.Scan(&name, &inv_c); err != nil {
			log.Println(err)
		}
		fmt.Println(4)
		if name == action.Data.Name && inv_c == action.Data.InviteCode {
			us = `INSERT INTO users_rooms(Rooms_id, User_id) VALUES (?, ?)`
			lg, err = db.Query(us, action.Data.ID, action.Data.User_ID)
			response, err := json.Marshal(Response{Action: "login", Success: true, ObjName: "room", Data: User{ID: action.Data.User_ID}})
			if err != nil {
				panic(err)
			}
			return response
		}
		if err != nil {
			panic(err)
		}
	}
	response, err := json.Marshal(Response{Action: "login", Success: false, ObjName: "room"})
	return response
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

func (action CreateRoom) Process() []byte { return nil }

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

func (action EditRoom) Process() []byte { return nil }

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

func (action DeleteRoom) Process() []byte {
	return nil
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

func (action ReadRoom) Process() []byte { return nil }

func (r Room) Print()        {}
func (r Room) GetID() string { return "" }
