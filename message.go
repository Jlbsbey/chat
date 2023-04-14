package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	Cont        Content   `json:"content"`
	Message_ID  uint64    `json:"message_id"`
	Room_ID     int       `json:"room_id"`
	Author      int       `json:"author_id"`
	Reply       uint64    `json:"reply_message_id"`
	SendTime    time.Time `json:"date"`
	IsChanged   bool      `json:"is_changed"`
	IsDeleted   bool      `json:"is_deleted"`
	IsForwarded bool      `json:"is_forwarded"`
	Username    string    `json:"author"`
}
type LoginMessage struct{}
type LogoutMessage struct{}
type CreateMessage struct {
	Data struct {
		Cont        Content   `json:"content"`
		Author      int       `json:"author_id"`
		Room_ID     int       `json:"room_id"`
		Reply       uint64    `json:"reply_message_id"`
		SendTime    time.Time `json:"date"`
		IsForwarded bool      `json:"is_forwarded"`
	} `json:"data"`
}
type EditMessage struct {
	M Message `json:"data"`
}
type ReadMessage struct {
	Data struct {
		Message_ID uint64 `json:"message_id"`
		Room_ID    int    `json:"room_id"`
	} `json:"data"`
}
type DeleteMessage struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (m Message) Login() DefinedAction {
	return &LoginMessage{}
}
func (action *LoginMessage) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action LoginMessage) Process() []byte {
	return nil
}
func (m Message) Logout() DefinedAction {
	return &LogoutMessage{}
}
func (action *LogoutMessage) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action LogoutMessage) Process() []byte {
	return nil
}

func (m Message) Create() DefinedAction {
	return &CreateMessage{}
}

func (action *CreateMessage) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action CreateMessage) Process() []byte {
	var MessageID uint64
	Attributes := []byte("000")
	MessageID = CreateMessageID()
	if action.Data.IsForwarded == true {
		Attributes[2] = '1'
	}
	MsgText := []byte(action.Data.Cont.Text)
	fmt.Println(action.Data.Cont.Text)
	fmt.Println(MsgText)
	us := `INSERT INTO messages (Message_ID, Author_ID, Room_ID, Text, Creation_time, Attributes, ReplyToMesID) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.ExecContext(context.Background(), us, MessageID, action.Data.Author, action.Data.Room_ID, MsgText, action.Data.SendTime, Attributes, action.Data.Reply)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(Response{Action: "create", Success: true, ObjName: "message", Data: Message{Message_ID: MessageID}})
	if err != nil {
		panic(err)
	}
	return response
}

func (m Message) Edit() DefinedAction {
	return &EditMessage{}
}

func (action *EditMessage) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action EditMessage) Process() []byte { return nil }

func (m Message) Delete() DefinedAction {
	return &DeleteRoom{}
}

func (action *DeleteMessage) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action DeleteMessage) Process() []byte {
	return nil
}

func (m Message) Read() DefinedAction {
	return &ReadMessage{}
}
func (action *ReadMessage) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadMessage) Process() []byte {
	us := `SELECT Author_ID, Room_ID, Text, Creation_time, Attributes, ReplyToMesID FROM messages WHERE Message_id = ? OR Room_ID = ?`

	rows, err := db.Query(us, action.Data.Message_ID, action.Data.Room_ID)
	if err != nil {
		panic(err)
	}
	var Author_id, Room_id int
	var Creation_time time.Time
	var temptime []uint8
	var Attributes, Text, Login, quer string
	var IsForwarded bool
	var ReplyToMesID uint64
	var Messages []Message
	var lg *sql.Rows
	for rows.Next() {
		err = rows.Scan(&Author_id, &Room_id, &Text, &temptime, &Attributes, &ReplyToMesID)
		fmt.Println(Text)
		if err != nil {
			panic(err)
		}
		datetimeStr := string(temptime)
		Creation_time, err = time.Parse("2006-01-02 15:04:05", datetimeStr)
		if err != nil {
			panic(err)
		}
		if Attributes[2] == '1' {
			IsForwarded = true
		} else {
			IsForwarded = false
		}
		quer = `SELECT Login FROM users WHERE ID = ?`
		lg, err = db.Query(quer, Author_id)
		if err != nil {
			panic(err)
		}
		for lg.Next() {
			err = lg.Scan(&Login)
			if err != nil {
				panic(err)
			}
			//append to array
			Messages = append(Messages, Message{Message_ID: action.Data.Message_ID, Author: Author_id, Room_ID: Room_id, Cont: Content{Text: Text}, SendTime: Creation_time, IsForwarded: IsForwarded, Reply: ReplyToMesID, Username: Login})

		}

	}
	if len(Messages) == 0 {
		response, err := json.Marshal(Response{Action: "read", Success: false, ObjName: "message", Error_message: "Message not found"})
		if err != nil {
			panic(err)
		}
		return response
	} else {
		response, err := json.Marshal(Response{Action: "read", Success: true, ObjName: "message", Data: Messages})
		if err != nil {
			panic(err)
		}
		return response
	}
}

func (m Message) Print()        {}
func (m Message) GetID() string { return "" }

func CreateMessageID() uint64 {
	var MessageID uint64
	var free bool
	free = false
	var max, min int
	max = 1000000000000000
	min = 1
	for free == false {
		rand.Seed(time.Now().UnixNano())
		MessageID = uint64(rand.Intn(max-min) + min)
		us := `SELECT Message_ID FROM messages WHERE Message_ID = ?`
		rows, err := db.Query(us, MessageID)
		if err != nil {
			panic(err)
		}
		if rows.Next() == false {
			free = true
		}
	}
	return MessageID
}
