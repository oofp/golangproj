package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

//go run github.com/99designs/gqlgen generate

import (
	"errors"
	"sync"

	"github.com/brozinov/gogql/graph/model"
)

var todos sync.Map
var todochannel chan *model.TodoEvent

func AddNewTodo(todo model.NewTodo) *model.Todo {
	new_id := GetNextId()
	user := model.User{ID: todo.UserID}
	newTodo := model.Todo{ID: new_id, Text: todo.Text, Done: false, User: &user}
	todos.Store(new_id, &newTodo)
	sendEvent(model.EventTypeToDoAdded, &newTodo)
	return &newTodo
}

func sendEvent(evType model.EventType, td *model.Todo) {
	if todochannel != nil {
		evType := evType
		ev := model.TodoEvent{EvType: &evType, Todo: td}
		todochannel <- &ev
	}
}

func RemoveTodo(id string) (*model.Todo, error) {
	deletedToDo, flExist := todos.LoadAndDelete(id)
	if flExist {
		td := deletedToDo.(*model.Todo)
		sendEvent(model.EventTypeToDoRemoved, td)
		return td, nil
	} else {
		return nil, errors.New("Todo not found")
	}
}

func SetTodoDone(id string, done *bool) (*model.Todo, error) {
	existingToDo, flExist := todos.Load(id)
	if flExist {
		td := existingToDo.(*model.Todo)
		td.Done = *done
		sendEvent(model.EventTypeToDoUpdated, td)
		return td, nil
	} else {
		return nil, errors.New("Todo not found")
	}
}

func GetAllTodos() [](*model.Todo) {
	//allTodos := make([](*model.Todo), 1)
	var allTodos [](*model.Todo)

	todos.Range(func(k, v interface{}) bool {
		allTodos = append(allTodos, v.(*model.Todo))
		return true
	})

	return allTodos
}

func createSubChannel() {
	todochannel = make(chan *model.TodoEvent)
}

func get_subchannel() (newchannel chan *model.TodoEvent) {
	//fmt.Println("get_subchannel invoked")
	createSubChannel()
	return todochannel
}
