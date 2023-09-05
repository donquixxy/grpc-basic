package main

import (
	"context"
	"fmt"
	"gokit-basic/common/config"
	m "gokit-basic/common/model"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localUser *m.ListUsers

func main() {

	initUser()

	srv := grpc.NewServer()

	reflection.Register(srv)

	m.RegisterUsersServer(srv, &UsersServer{})

	listener, err := net.Listen("tcp", config.SERVICE_USER_PORT)

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	log.Println(listener.Addr())

	go func() {
		if err = srv.Serve(listener); err != nil {
			log.Fatalf("failed to serve %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
}

func initUser() {

	localUser = new(m.ListUsers)
	localUser.List = make([]*m.UserServices, 0)
}

type UsersServer struct {
	m.UnimplementedUsersServer
}

func (s *UsersServer) CreateUser(ctx context.Context, v *m.UserServices) (*m.UserServices, error) {

	log.Println("called func")

	if v.Name == "" {
		log.Println("error here")
		return nil, status.Error(codes.Unknown, "name cannot be empty")
	}

	localUser.List = append(localUser.List, v)

	log.Println("User created :", localUser.String())

	return v, nil
}

func (s *UsersServer) GetListUser(ctx context.Context, v *emptypb.Empty) (*m.ListUsers, error) {
	return &m.ListUsers{
		List: localUser.List,
	}, nil
}

func (s *UsersServer) GetByName(ctx context.Context, v *m.ByName) (*m.UserServices, error) {

	for _, item := range localUser.List {
		if item.Name == v.Name {
			return item, nil
		}
	}
	return nil, status.Error(codes.NotFound, fmt.Sprintf("User %s not found", v.Name))
}
