package chatgo

import (
	"github.com/labstack/echo"
)

type Chat struct {
	scenarios map[string]Scenario
}

func NewChat() Chat {
	return Chat{make(map[string]Scenario)}
}

// Set scenarios[event] as given scenario if event is not in key.
func (chat Chat) Set(event string, scenario Scenario) {
	if _, ok := chat.scenarios[event]; !ok {
		chat.scenarios[event] = scenario
	}
}

func (chat Chat) getKeyboard(c echo.Context) error {
	return c.JSON(200, RunScenario(chat.scenarios["get_keyboard"], nil))
}

func (chat Chat) postMessage(c echo.Context) error {
	req := new(struct {
		UserKey string `json:"user_key"`
		Type    string `json:"type"`
		Content string `json:"content"`
	})

	if err := c.Bind(req); err != nil {
		return err
	}

	var input Object

	switch req.Type {
	case "text":
		input = Text(req.Content)
	case "photo":
		input = Photo{Url: req.Content}
	default:
		return echo.NewHTTPError(501, "invalid request format")
	}

	return c.JSON(200, RunScenario(chat.scenarios["reply_message"], input))
}

func (chat Chat) postFriend(c echo.Context) error {
	req := new(struct {
		UserKey string `json:"user_key"`
	})

	if err := c.Bind(req); err != nil {
		return err
	}

	RunScenario(
		chat.scenarios["add_user"],
		Text(req.UserKey),
	)

	return c.JSON(200, &struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Comment string `json:"comment"`
	}{Code: 0, Message: "SUCCESS", Comment: "정상 응답"})
}

func (chat Chat) deleteFriend(c echo.Context) error {
	userKey := c.Param("user_key")

	RunScenario(
		chat.scenarios["delete_user"],
		Text(userKey),
	)

	return c.JSON(200, &struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Comment string `json:"comment"`
	}{Code: 0, Message: "SUCCESS", Comment: "정상 응답"})
}

func (chat Chat) deleteRoom(c echo.Context) error {
	userKey := c.Param("user_key")

	RunScenario(
		chat.scenarios["delete_room"],
		Text(userKey),
	)

	return c.JSON(200, &struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Comment string `json:"comment"`
	}{Code: 0, Message: "SUCCESS", Comment: "정상 응답"})
}

// New returns new *echo.Echo and error to run.
func (chat Chat) New() *echo.Echo {
	e := echo.New()

	e.GET("/keyboard", chat.getKeyboard)
	e.POST("/message", chat.postMessage)
	e.POST("/friend", chat.postFriend)
	e.DELETE("/friend/:user_key", chat.deleteFriend)
	e.DELETE("/chat_room/:user_key", chat.deleteRoom)

	return e
}
