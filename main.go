package main

import (
	"context"
	"net"
    "log"
	"todoService/pb"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedTodoServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*
	pb.CreateTodoResponse, error) {
	return &pb.CreateTodoResponse{}, nil
}

func (s *server) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*
	pb.DeleteTodoResponse, error) {
    return &pb.DeleteTodoResponse{}, nil
}

func (s *server) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*
	pb.UpdateTodoResponse, error) {
	return &pb.UpdateTodoResponse{}, nil
}

func (s *server) ListTodos(ctx context.Context, req *pb.ListTodosRequest) (*
	pb.ListTodosResponse, error) {
	return &pb.ListTodosResponse{}, nil
}

func (s *server) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*
	pb.GetTodoResponse, error) {
	return &pb.GetTodoResponse{}, nil
}