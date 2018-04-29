package chatgo

import (
	"github.com/labstack/echo"
)

type Chat struct {
	scenarios map[string]Scenario
}

// New returns new *echo.Echo and error to run.
func (chat Chat) New() *echo.Echo {
	e := echo.New()

	e.GET("/keyboard", getKeyboard)
	e.POST("/message", postMessage)
	e.POST("/friend", postFriend)
	e.DELETE("/friend/:user_key", deleteFriend)
	e.DELETE("/chat_room/:user_key", deleteRoom)

	return e
}

// Add sets scenarios[event] as given scenario if event is not in key.
func (chat Chat) Add(event string, scenario Scenario) {
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
		input = Photo{req.Content}
	default:
		return echo.NewHTTPError(501, "invalid request format")
	}

	return c.JSON(200, RunScenario(chat.scenarios["reply_message"], input))
}

func (chat Chat) postFriend(c echo.Context) error {
	req := new(struct {
		UserKey string `json:"user_key"`
	})

	if err := c.Bind(u); err != nil {
		return err
	}

	return c.String(200, string(
		RunScenario(
			chat.scenarios["add_user"],
			Text(req.UserKey),
		).(Text)))
}

func (chat Chat) deleteFriend(c echo.Context) error {
	userKey := c.Param("user_key")

	return c.String(200, string(
		RunScenario(chat.scenarios["delete_user"],
			Text(userKey),
		).(Text)))
}

func (chat Chat) deleteRoom(c echo.Context) error {
	userKey := c.Param("user_key")

	return c.String(200, string(
		RunScenario(chat.scenarios["delete_room"],
			Text(userKey),
		).(Text)))
}
