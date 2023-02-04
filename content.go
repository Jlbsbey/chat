package main

import "image"

type Content struct {
	Text  string      `json:"text"`
	Image image.Image `json:"image"`
}
