package main

import "time"

type Message struct {
	Cont      Content
	Author    uint64    `json:"author_id"`
	Reply     uint64    `json:"reply_message_id"`
	SendTime  time.Time `json:"date"`
	IsChanged bool      `json:"is_changed"`
	IsDeleted bool      `json:"is_deleted"`
}
