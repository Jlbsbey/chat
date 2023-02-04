package main

type User struct {
	name     string          `json:"name"`
	password string          `json:"password"`
	ID       uint64          `json:"id"`
	email    string          `json:"email"`
	room     map[uint64]bool `json:"room_id"`
}
