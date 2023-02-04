package main

import "time"

type Message struct {
	attribute uint32
	author    uint64
	reply     uint64
	sendTime  time.Time
}
