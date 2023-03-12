package main

type Request struct {
	Object string      `json:"object"`
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
