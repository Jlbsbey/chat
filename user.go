package chat

type User struct {
	name     string
	password string
	ID       uint64
	email    string
	room     map[uint64]bool
}
