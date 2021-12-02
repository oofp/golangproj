package main

import (
	"context"
	"errors"
	"fmt"
	"io"
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

	todoupd, err := s.addNewToDo(newToDo)
	if err != nil {
		return nil, err
	}
	return todoupd.Todo, nil
}

func (s *TodoGrpcServer) GetAllToDos(_ *pb.NoParams, resp pb.TodoGrpc_GetAllToDosServer) error {
	s.todos.Range(func(k, v interface{}) bool {

		resp.Send(v.(*pb.ToDo))
		return true
	})

	return nil
}

func (s *TodoGrpcServer) MonitorToDos(_ *pb.NoParams, mon pb.TodoGrpc_MonitorToDosServer) error {
	return nil
}

func (s *TodoGrpcServer) ApplyOps(ops pb.TodoGrpc_ApplyOpsServer) error {
	log.Println("Entered ApplyOps")

	for {
		op, err := ops.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("ApplyOps terminated by client")
			} else {
				log.Printf("ApplyOps error:%v", err)
			}

			break
		}

		log.Printf("ApplyOps operation:%v", op)
		{
			todo, err := s.handleOp(op)
			log.Printf("ApplyOps operation result:%v %v", todo, err)
			if err == nil {
				log.Printf("ApplyOps todo update:%v", todo)
				ops.Send(todo)
			}
		}
	}
	return nil
}

func (s *TodoGrpcServer) handleOp(op *pb.ToDoOp) (*pb.ToDoUpdate, error) {

	switch opr := op.GetOp().(type) {
	case *pb.ToDoOp_Add:
		newToDo := opr.Add.GetTodo()
		return s.addNewToDo(newToDo)
	case *pb.ToDoOp_Upd:
		upd := opr.Upd
		return s.updateToDo(upd.GetTodoID(), upd.GetDone())
	case *pb.ToDoOp_Del:
		delToDo := opr.Del
		return s.deleteToDo(delToDo.GetTodoID())
	}

	return nil, errors.New("Unexpected")
}

func (s *TodoGrpcServer) addNewToDo(newToDo *pb.NewToDo) (*pb.ToDoUpdate, error) {
	td_id := fmt.Sprint(rand.Intn(10000))
	todo := pb.ToDo{Task: newToDo.GetTask(), Done: false, UserID: newToDo.UserID, TodoID: td_id}
	s.todos.Store(todo.TodoID, &todo)

	return &pb.ToDoUpdate{UpdateType: pb.ToDoUpdateType_Added, Todo: &todo}, nil
}

func (s *TodoGrpcServer) updateToDo(todoId string, done bool) (*pb.ToDoUpdate, error) {
	todo, ok := s.todos.Load(todoId)

	if ok {
		todoptr := todo.(*pb.ToDo)
		todoptr.Done = done
		upd := pb.ToDoUpdate{UpdateType: pb.ToDoUpdateType_Updated, Todo: todoptr}
		return &upd, nil
	} else {
		return nil, errors.New("todo not found")
	}
}

func (s *TodoGrpcServer) deleteToDo(todoId string) (*pb.ToDoUpdate, error) {
	todo, ok := s.todos.LoadAndDelete(todoId)

	if ok {
		todoptr := todo.(*pb.ToDo)
		return &pb.ToDoUpdate{UpdateType: pb.ToDoUpdateType_Deleted, Todo: todoptr}, nil
	} else {
		return nil, errors.New("todo not found")
	}
}

func (s *TodoGrpcServer) KeepAlive(ka pb.TodoGrpc_KeepAliveServer) error {

	for {
		ping, err := ka.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Keep alive terminated by client")
			} else {
				log.Printf("Ping error:%v", err)
			}

			break
		}

		msgID := ping.GetMsgID()

		pong := pb.Pong{MsgID: msgID}

		ka.Send(&pong)

		if msgID == 0 {
			log.Printf("Keep alive 0, exiting loop")
			break
		}
	}

	return nil
}

func (s TodoGrpcServer) Kill(ctx context.Context, _ *pb.NoParams) (*pb.NoParams, error) {
	log.Fatal("Kill received")
	return &pb.NoParams{}, nil
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
