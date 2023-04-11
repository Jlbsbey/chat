package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type Room struct {
	Name string `json:"name"`
	//Messages   []Message `json:"messages"`
	Room_ID    int `json:"room_id"`
	InviteCode int `json:"inv_code"`
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
	Data struct {
		Name    string `json:"name"`
		User_ID int    `json:"user_id"`
	} `json:"data"`
}
type EditRoom struct {
	R Room `json:"data"`
}
type ReadRoom struct {
	Data struct {
		User_ID int `json:"user_id"`
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
	us := `SELECT ID, Room_name, Invite_code FROM rooms WHERE Room_name = ? OR Invite_code = ?`
	fmt.Println(action.Data.Name)
	fmt.Println(action.Data.InviteCode)
	fmt.Println(action.Data.User_ID)
	lg, err := db.Query(us, action.Data.Name, action.Data.InviteCode)
	if err != nil {
		panic(err)
	}
	var name string
	var inv_c, ID int
	for lg.Next() {
		if err = lg.Scan(&ID, &name, &inv_c); err != nil {
			log.Println(err)
		}
		if name == action.Data.Name && inv_c == action.Data.InviteCode {
			us = `INSERT INTO users_rooms(Room_id, User_id) VALUES (?, ?)`
			_, err = db.ExecContext(context.Background(), us, ID, action.Data.User_ID)
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

func (action CreateRoom) Process() []byte {
	us := `INSERT INTO rooms(Room_name) VALUES (?)`
	fmt.Println(action.Data.Name)
	fmt.Println(action.Data.User_ID)
	_, err := db.ExecContext(context.Background(), us, action.Data.Name)
	us = `SELECT ID, Room_name FROM rooms WHERE Room_name = ?`
	lg, err := db.Query(us, action.Data.Name)
	if err != nil {
		panic(err)
	}
	var name string
	var ID int
	for lg.Next() {
		if err = lg.Scan(&ID, &name); err != nil {
			log.Println(err)
		}
	}
	us = `INSERT INTO users_rooms(Room_id, User_id) VALUES (?, ?)`
	_, err = db.ExecContext(context.Background(), us, ID, action.Data.User_ID)
	response, err := json.Marshal(Response{Action: "login", Success: true, ObjName: "room"})
	if err != nil {
		panic(err)
	}
	return response
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

func (action ReadRoom) Process() []byte {
	var Rooms []Room
	us := `SELECT Room_id, Attributes FROM users_rooms WHERE User_id = ?`
	lg, err := db.Query(us, action.Data.User_ID)
	if err != nil {
		panic(err)
	}
	var attributes string
	var room_id int
	for lg.Next() {
		if err = lg.Scan(&room_id, &attributes); err != nil {
			log.Println(err)
		}
		if attributes[8] == '0' {
			us = `SELECT Room_name FROM rooms WHERE ID = ?`
			lg, err := db.Query(us, room_id)
			if err != nil {
				panic(err)
			}
			var name string
			for lg.Next() {
				if err = lg.Scan(&name); err != nil {
					log.Println(err)
				}
				Rooms = append(Rooms, Room{Name: name, Room_ID: room_id})
			}
		}
	}
	response, err := json.Marshal(Response{Action: "read", ObjName: "room", Data: Rooms})
	if err != nil {
		panic(err)
	}
	return response

}

func (r Room) Print()        {}
func (r Room) GetID() string { return "" }
