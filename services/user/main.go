package main

import (
	"context"
	"gokit-basic/common/config"
	m "gokit-basic/common/model"
	"gokit-basic/repository/postgres"
	"gokit-basic/services/user/domain"
	"gokit-basic/services/user/repo"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var vdtor *validator.Validate

func initValidator() {
	vdtor = validator.New()
}

var localUser *m.ListUsers

func main() {

	config := config.InitConfig()
	db := postgres.InitDatabase(config)
	repo := repo.NewUserRepo(db)

	srv := grpc.NewServer()
	initValidator()
	reflection.Register(srv)

	m.RegisterUsersServer(srv, &UsersServer{
		userRepository: repo,
	})

	listener, err := net.Listen("tcp", config.ServerConf.SERVICE_USER_PORT)

	log.Println(config.ServerConf.SERVICE_USER_PORT)

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

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

type UsersServer struct {
	userRepository repo.UserRepo
	m.UnimplementedUsersServer
}

func (s *UsersServer) CreateUser(ctx context.Context, v *m.SingleUser) (*m.SingleUser, error) {

	log.Println("called func")

	err := vdtor.Var(v.Name, "required")

	if err != nil {
		log.Printf("name is empty %v", err)
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}

	if err = vdtor.Var(v.Phone, "required"); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Phone is required")
	}

	if err = vdtor.Var(v.Age, "required"); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Age is required")
	}

	_, err = s.userRepository.CreateUser(domain.ToUserDomainMapper(v))

	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	return v, nil
}

func (s *UsersServer) GetListUser(ctx context.Context, v *emptypb.Empty) (*m.ListUsers, error) {
	lists := new(m.ListUsers)

	listItem := s.userRepository.GetListUser()

	if len(listItem) == 0 {
		return nil, status.Error(codes.NotFound, "no data found")
	}

	for _, i := range listItem {
		lists.List = append(lists.List, i.ToUserProtoMappter())
	}

	return lists, nil
}
