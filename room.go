package chat

type Room struct {
	name     string
	Messages []Message
	id       uint64
	users    map[uint64]uint32
}
