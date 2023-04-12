package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, X-Auth-Token")
	if req.Method == "OPTIONS" {
		w.WriteHeader(204)
	} else if req.Method == "GET" {

	} else if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		w.Write(RecAction(data))
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	cfg()
	http.HandleFunc("/", Handler)

	err := http.ListenAndServe(":8080", nil)
	panic(err)
}

func RecAction(text []byte) []byte {
	logined = false
	var action Action
	err := json.Unmarshal(text, &action)
	if err != nil {
		panic(err)
	}
	var obj GeneralObject
	switch action.ObjName {
	case "user":
		obj = &User{}
	case "room":
		obj = &Room{}
	case "message":
		obj = &Message{}
	default:
		fmt.Println("Unknown object", action.ObjName)

	}
	var defact DefinedAction
	switch action.Action {
	case "register":
		defact = obj.Create()
	case "edit":
		//defact = obj.Edit()
	case "delete":
		//defact = obj.Delete()
	case "read":
		defact = obj.Read()
	case "login":
		if action.ObjName == "message" {
			return []byte("")
		}
		defact = obj.Login()
	case "logout":
		if action.ObjName == "message" {
			return []byte("")
		}
		defact = obj.Logout()
	default:
		panic("Unknown action")
	}
	defact.GetFromJSON(text)
	return defact.Process()

}
