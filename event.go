package chatgo

type Event struct {
	Type   string
	Method string
	Info   map[string]Object
}

func NewEvent_AddUser(id string) Event {
	return Event{
		Type:   "user",
		Method: "add",
		Info:   map[string]Object{"id": Text{id}},
	}
}

func NewEvent_DeleteUser(id string) Event {
	return Event{
		Type:   "user",
		Method: "delete",
		Info:   map[string]Object{"id": Text{id}},
	}
}

func NewEvent_CallKeyboard() Event {
	return Event{
		Type:   "keyboard",
		Method: "call",
		Info:   map[string]Object{},
	}
}

func NewEvent_RequestMessage(id, t, content string) Event {
	var info Object
	switch t {
	case "text":
		info = Text{content}
	case "photo":
		info = Photo{content}
	}

	return Event{
		Type:   "message",
		Method: "request",
		Info: map[string]Object{"content": info,
			"id":   Text{id},
			"type": Text{t},
		},
	}
}
