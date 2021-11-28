package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	pb "oofp/todogrpc/todoproto"
)

const (
	port = ":9098"
)

type TodoGrpcServer struct {
	pb.UnimplementedTodoGrpcServer
}

func (s *TodoGrpcServer) CreateNewToDo(ctx context.Context, newToDo *pb.NewToDo) (*pb.ToDo, error) {
	log.Printf("Received: %v", newToDo.GetTask())
	td_id := fmt.Sprint(rand.Intn(10000))

	return &pb.ToDo{Task: newToDo.GetTask(), Done: false, UserID: newToDo.UserID, TodoID: td_id}, nil
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
