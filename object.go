package chatgo

import "encoding/json"

type Object interface {
	Get() interface{}
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

/*
Kakao Plusfriend
*/
type Keyboard struct {
	Buttons []string
}

func (k Keyboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    string   `json:"type"`
		Buttons []string `json:"buttons,omitempty"`
	}{Type: "buttons", Buttons: k.Buttons})
}

type Message struct {
	Text  string `json:"text,omitempty"`
	Photo Photo  `json:"photo,omitempty"`
}

type Photo struct {
	Url    string
	Width  int
	Height int
}

func (p Photo) MarshalJSON() ([]byte, error) {
	width := p.Width
	height := p.Height
	if width == 0 {
		width = 720
	}
	if height == 0 {
		height = 630
	}
	return json.Marshal(&struct {
		Url    string `json:"url,omitempty"`
		Width  int    `json:"width,omitempty"`
		Height int    `json:"height,omitempty"`
	}{Url: p.Url, Width: width})
}
