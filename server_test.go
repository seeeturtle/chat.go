package chatgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

type Comparison func(*httptest.ResponseRecorder) string

type T struct {
	handlerFunc echo.HandlerFunc
	t           *testing.T
	req         *http.Request
	rec         *httptest.ResponseRecorder
	path        string
	comparisons map[string]Comparison
}

func testHandler(t T) {
	e := echo.New()
	c := e.NewContext(t.req, t.rec)
	c.SetPath(t.path)

	if err := t.handlerFunc(c); err != nil {
		t.t.Errorf("got error while calling handlerFunc: %v\n", err)
	}
	for name, comparison := range t.comparisons {
		t.t.Run(
			fmt.Sprintf("%s_%s", t.t.Name(), name),
			func(test *testing.T) {
				if err := comparison(t.rec); err != "" {
					test.Error(err)
				}
			},
		)
	}
}

func indentedJSON(s string) (string, error) {
	var output bytes.Buffer
	err := json.Indent(&output, []byte(s), "", "\t")
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

func compareCode(code int) Comparison {
	return func(rec *httptest.ResponseRecorder) string {
		if rec.Code != code {
			return fmt.Sprintf("expected %d but got %d", code, rec.Code)
		}
		return ""
	}
}

func compareJSON(expected string) Comparison {
	return func(rec *httptest.ResponseRecorder) string {
		if rec.Body.String() != expected {
			expected, err := indentedJSON(expected)
			if err != nil {
				return fmt.Sprintf("got error while indenting JSON: %s", err)
			}
			got, err := indentedJSON(rec.Body.String())
			if err != nil {
				return fmt.Sprintf("got error while indenting JSON: %s", err)
			}
			return fmt.Sprintf("expected:\n%s\ngot:\n%s", expected, got)
		}
		return ""
	}
}

func TestGetKeyboard(t *testing.T) {
	var getKeyboardScenario CondScenario
	getKeyboardScenario.Add(
		func(o Object) bool {
			return true
		},
		func(o Object) (Scenario, Object) {
			return nil, Keyboard{[]string{"hello", "seeeturtle"}}
		},
	)

	chat := NewChat()
	chat.Add("get_keyboard", getKeyboardScenario)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	expectedJSON := `{"type":"buttons","buttons":["hello","seeeturtle"]}`

	testHandler(
		T{
			handlerFunc: chat.getKeyboard,
			t:           t,
			req:         req,
			rec:         httptest.NewRecorder(),
			path:        "/keyboard",
			comparisons: map[string]Comparison{
				"Code": compareCode(200),
				"JSON": compareJSON(expectedJSON),
			},
		},
	)
}

func TestPostMessage(t *testing.T) {
	var postMessageScenario CondScenario
	postMessageScenario.Add(
		func(o Object) bool { return true },
		func(o Object) (Scenario, Object) {
			switch v := o.(type) {
			case Text:
				return nil, Message{Text: string(v)}
			default:
				return nil, nil
			}
		},
	)

	chat := NewChat()
	chat.Add("reply_message", postMessageScenario)

	req := httptest.NewRequest(
		"POST",
		"/",
		strings.NewReader(`{"user_key":"encryptedUserKey","type":"text","content":"Hello World"}`),
	)
	req.Header.Set("Content-Type", "application/json")

	expectedJSON := `{"message":{"text":"Hello World"}}`

	testHandler(
		T{
			handlerFunc: chat.postMessage,
			t:           t,
			req:         req,
			rec:         httptest.NewRecorder(),
			path:        "/message",
			comparisons: map[string]Comparison{
				"Code": compareCode(200),
				"JSON": compareJSON(expectedJSON),
			},
		},
	)
}

func TestPostFriend(t *testing.T) {
	var postFriendScenario CondScenario
	postFriendScenario.Add(
		func(o Object) bool { return true },
		func(Object) (Scenario, Object) {
			return nil, nil
		},
	)

	chat := NewChat()
	chat.Add("add_user", postFriendScenario)

	req := httptest.NewRequest("POST",
		"/",
		strings.NewReader(`{"user_key":"user_name"}`),
	)
	req.Header.Set("Content-Type", "application/json")

	expectedJSON := `{"code":0,"message":"SUCCESS","comment":"정상 응답"}`

	testHandler(
		T{
			handlerFunc: chat.postFriend,
			t:           t,
			req:         req,
			rec:         httptest.NewRecorder(),
			path:        "/friend",
			comparisons: map[string]Comparison{
				"Code": compareCode(200),
				"JSON": compareJSON(expectedJSON),
			},
		},
	)
}

func TestDeleteFriend(t *testing.T) {
	var deleteFriendScenario CondScenario
	deleteFriendScenario.Add(
		func(o Object) bool { return true },
		func(o Object) (Scenario, Object) { return nil, nil },
	)

	chat := NewChat()
	chat.Add("delete_user", deleteFriendScenario)

	req := httptest.NewRequest(
		"DELETE",
		"/",
		strings.NewReader(`{"user_key":"user_name"}`),
	)
	req.Header.Set("Content-Type", "application/json")

	expectedJSON := `{"code":0,"message":"SUCCESS","comment":"정상 응답"}`

	testHandler(
		T{
			handlerFunc: chat.deleteFriend,
			t:           t,
			req:         req,
			rec:         httptest.NewRecorder(),
			path:        "/friend/:user_key",
			comparisons: map[string]Comparison{
				"Code": compareCode(200),
				"JSON": compareJSON(expectedJSON),
			},
		},
	)
}

func TestDeleteRoom(t *testing.T) {
	var deleteRoomScenario CondScenario
	deleteRoomScenario.Add(
		func(o Object) bool { return true },
		func(o Object) (Scenario, Object) { return nil, nil },
	)

	chat := NewChat()
	chat.Add("delete_room", deleteRoomScenario)

	req := httptest.NewRequest(
		"DELETE",
		"/",
		strings.NewReader(`{"user_key":"user_name"}`),
	)
	req.Header.Set("Content-Type", "application/json")

	expectedJSON := `{"code":0,"message":"SUCCESS","comment":"정상 응답"}`

	testHandler(
		T{
			handlerFunc: chat.deleteRoom,
			t:           t,
			req:         req,
			rec:         httptest.NewRecorder(),
			path:        "/room/:user_key",
			comparisons: map[string]Comparison{
				"Code": compareCode(200),
				"JSON": compareJSON(expectedJSON),
			},
		},
	)
}
