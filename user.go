package main

type User struct {
	Name     string          `json:"name"`
	Password string          `json:"password"`
	ID       uint64          `json:"id"`
	Email    string          `json:"email"`
	Room     map[uint64]bool `json:"room_id"`
}
