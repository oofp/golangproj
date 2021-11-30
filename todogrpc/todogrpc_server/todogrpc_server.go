package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	pb "oofp/todogrpc/todoproto"
	"sync"

	"google.golang.org/grpc"
)

const (
	port = ":9098"
)

type TodoGrpcServer struct {
	pb.UnimplementedTodoGrpcServer
	todos sync.Map
}

func (s *TodoGrpcServer) CreateNewToDo(ctx context.Context, newToDo *pb.NewToDo) (*pb.ToDo, error) {
	log.Printf("Received: %v", newToDo.GetTask())
	td_id := fmt.Sprint(rand.Intn(10000))

	todo := pb.ToDo{Task: newToDo.GetTask(), Done: false, UserID: newToDo.UserID, TodoID: td_id}
	s.todos.Store(todo.TodoID, &todo)

	return &todo, nil
}

func (s *TodoGrpcServer) GetAllToDos(_ *pb.NoParams, resp pb.TodoGrpc_GetAllToDosServer) error {
	s.todos.Range(func(k, v interface{}) bool {

		resp.Send(v.(*pb.ToDo))
		return true
	})

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoGrpcServer(s, &TodoGrpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
