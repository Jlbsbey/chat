package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Room struct {
	Name       string    `json:"name"`
	Messages   []Message `json:"messages"`
	ID         string    `json:"id"`
	InviteCode int       `json:"inv_code"`
}

// скопиировать name id invite code сюда и использовать этот стракт
type LoginRoom struct {
	Data struct {
		R       Room `json:"data"`
		User_ID int  `json:"user_id"`
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
	us := `SELECT Name, Invite_code FROM rooms WHERE Name = ? AND Invite_code = ?`

	lg, err := db.Query(us, action.R.Name, action.R.InviteCode)
	if err != nil {
		panic(err)
	}
	var name string
	var inv_c int
	var ID int
	for lg.Next() {
		if err = lg.Scan(&name, &inv_c); err != nil {
			log.Println(err)
		}
		if name == action.R.Name && inv_c == action.R.InviteCode {
			ses_id := CreateSession(ID)
			response, err := json.Marshal(Response{Session_ID: ses_id, Action: "login", Success: true, ObjName: "user", User_ID: ID})
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

/*
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

func (action CreateRoom) Process() []byte{}

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

func (action EditRoom) Process() []byte {}

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

func (action DeleteRoom) Process()[]byte {}

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

func (action ReadRoom) Process()[]byte {}
*/
func (r Room) Print() {
	fmt.Println("Name", r.Name, ", ID", r.ID, ", Messages", r.Messages, ", Users", r.Users)
}
func (r Room) GetID() string {
	return r.ID
}
