package main

import (
	"context"
	"log"
	"net"

	"github.com/kapustkin/go_calendar_grpc/calendarpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type calendarServer struct {
}

func (c *calendarServer) Get(ctx context.Context, req *calendarpb.GetRequest) (*calendarpb.EventResponse, error) {
	return &calendarpb.EventResponse{Event: &calendarpb.Event{Uuid: "123123", Message: "Test 1"}}, nil
}

func (c *calendarServer) GetAll(ctx context.Context, req *calendarpb.GetAllRequest) (*calendarpb.AllEventResponse, error) {
	return &calendarpb.AllEventResponse{Events: []*calendarpb.Event{&calendarpb.Event{Uuid: "123123", Message: "Test 1"}, &calendarpb.Event{Uuid: "666", Message: "Test 2"}}}, nil
}

func (c *calendarServer) Add(ctx context.Context, req *calendarpb.AddRequest) (*calendarpb.OperationStatusResponse, error) {
	return &calendarpb.OperationStatusResponse{Sucess: true}, nil
}

func (c *calendarServer) Edit(ctx context.Context, req *calendarpb.EditRequest) (*calendarpb.OperationStatusResponse, error) {
	return &calendarpb.OperationStatusResponse{Sucess: false}, nil
}

func (c *calendarServer) Remove(ctx context.Context, req *calendarpb.RemoveRequst) (*calendarpb.OperationStatusResponse, error) {
	return &calendarpb.OperationStatusResponse{Sucess: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5900")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	calendarpb.RegisterCalendarEventsServer(grpcServer, &calendarServer{})
	grpcServer.Serve(lis)
}
