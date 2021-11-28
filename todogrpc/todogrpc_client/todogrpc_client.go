package main

import (
	"context"
	"log"
	pb "oofp/todogrpc/todoproto"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:9098"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("fatal failed to connect:%v", err)
	}
	defer conn.Close()

	c := pb.NewTodoGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_todos = make(map[int32]string)
	new_todos[1] = "Make breakfast"
	new_todos[2] = "Eat breakfast"

	for user_id, task := range new_todos {
		new_todo := pb.NewToDo{Task: task, UserID: user_id}
		r, err := c.CreateNewToDo(ctx, &new_todo)
		if err != nil {
			log.Printf("failed to create todo:%v", err)
		}
		log_todo(r)

	}

}

func log_todo(toDo *pb.ToDo) {
	log.Printf(`
	Todo details:
		id: %v
		task: %v
		userId: %v
		done: %v`,
		toDo.TodoID, toDo.Task, toDo.UserID, toDo.Done)
}
