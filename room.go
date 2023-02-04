package main

type Room struct {
	Name     string            `json:"name"`
	Messages []Message         `json:"messages"`
	Id       uint64            `json:"id"`
	Users    map[uint64]uint32 `json:"allowedUsersID"`
}
