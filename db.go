package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var FirstFreeID int = 1

type DB []GeneralObject

func (db DB) GetIndex(id string) int {
	for i, v := range db {
		if v.GetID() == id {
			return i
		}
	}
	return -1
}

func (db *DB) UseAction(path string) {
	text, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	var action Action
	err = json.Unmarshal(text, &action)
	if err != nil {
		fmt.Println(err)
		return
	}
	var obj GeneralObject
	switch action.ObjName {
	case "user":
		obj = &User{}
	case "room":
		obj = &Room{}
	default:
		fmt.Println("Unknown object", action.ObjName)
	}
	var toDo DefinedAction
	switch action.Action {
	case "create":
		toDo = obj.Create()
	case "edit":
		toDo = obj.Edit()
	case "delete":
		toDo = obj.Delete()
	case "read":
		toDo = obj.Read()
	default:
		fmt.Println("Unknown action", action.Action)
		return
	}
	toDo.GetFromJSON(text)
	toDo.Process(db)
}
