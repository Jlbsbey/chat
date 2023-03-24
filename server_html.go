package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == "GET" {

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
	/*case "room":
	obj = &Room{}*/
	default:
		fmt.Println("Unknown object", action.ObjName)
	}
	var defact DefinedAction
	switch action.Action {
	case "create":
		defact = obj.Create()
	case "edit":
		//defact = obj.Edit()
	case "delete":
		//defact = obj.Delete()
	case "read":
		//defact = obj.Read()
	case "login":
		defact = obj.Login()
	default:
		panic("Unknown action")
	}
	defact.GetFromJSON(text)
	return defact.Process()

}
