package main

type Action struct {
	Action  string      `json:"action"`
	ObjName string      `json:"object"`
	Data    interface{} `json:"data"`
}

type Response struct {
	Session_ID uint64 `json:"session_id"`
	ObjName    string `json:"object"`
	Action     string `json:"action"`
	Success    bool   `json:"success"`
	User_ID    int    `json:"user_id"`
}
