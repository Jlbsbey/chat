package main

type Action struct {
	Action  string      `json:"action"`
	ObjName string      `json:"object"`
	Data    interface{} `json:"data"`
}
