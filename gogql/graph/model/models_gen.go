// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type TodoEvent struct {
	EvType *EventType `json:"evType"`
	Todo   *Todo      `json:"todo"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EventType string

const (
	EventTypeToDoAdded   EventType = "ToDoAdded"
	EventTypeToDoRemoved EventType = "ToDoRemoved"
	EventTypeToDoUpdated EventType = "ToDoUpdated"
)

var AllEventType = []EventType{
	EventTypeToDoAdded,
	EventTypeToDoRemoved,
	EventTypeToDoUpdated,
}

func (e EventType) IsValid() bool {
	switch e {
	case EventTypeToDoAdded, EventTypeToDoRemoved, EventTypeToDoUpdated:
		return true
	}
	return false
}

func (e EventType) String() string {
	return string(e)
}

func (e *EventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventType", str)
	}
	return nil
}

func (e EventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
