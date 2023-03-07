package main

import (
	"context"
	"reflect"
	"testing"
	"todoService/pb"
)

func Test_server_CreateTodo(t *testing.T) {
	type fields struct {
		UnimplementedTodoServiceServer pb.UnimplementedTodoServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CreateTodoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateTodoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedTodoServiceServer: tt.fields.UnimplementedTodoServiceServer,
			}
			got, err := s.CreateTodo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.CreateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_DeleteTodo(t *testing.T) {
	type fields struct {
		UnimplementedTodoServiceServer pb.UnimplementedTodoServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.DeleteTodoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DeleteTodoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedTodoServiceServer: tt.fields.UnimplementedTodoServiceServer,
			}
			got, err := s.DeleteTodo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.DeleteTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_UpdateTodo(t *testing.T) {
	type fields struct {
		UnimplementedTodoServiceServer pb.UnimplementedTodoServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.UpdateTodoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UpdateTodoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedTodoServiceServer: tt.fields.UnimplementedTodoServiceServer,
			}
			got, err := s.UpdateTodo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_ListTodos(t *testing.T) {
	type fields struct {
		UnimplementedTodoServiceServer pb.UnimplementedTodoServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.ListTodosRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ListTodosResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedTodoServiceServer: tt.fields.UnimplementedTodoServiceServer,
			}
			got, err := s.ListTodos(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.ListTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.ListTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_GetTodo(t *testing.T) {
	type fields struct {
		UnimplementedTodoServiceServer pb.UnimplementedTodoServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.GetTodoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetTodoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedTodoServiceServer: tt.fields.UnimplementedTodoServiceServer,
			}
			got, err := s.GetTodo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
