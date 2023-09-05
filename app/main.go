package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"gokit-basic/common/config"
	m "gokit-basic/common/model"

	"github.com/goombaio/namegenerator"
	"google.golang.org/grpc"
)

func main() {

	se := userService()

	v, err := se.CreateUser(context.Background(), &m.UserServices{
		Name:  NameGenerator(),
		Phone: "01231",
		Age:   "123",
	})

	if err != nil {
		log.Fatalf("failed to create %v", err)
	}

	log.Println("Created user :", v)
}

func userService() m.UsersClient {

	conn, err := grpc.Dial(config.SERVICE_USER_PORT, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect server rpc %v", err)
	}

	return m.NewUsersClient(conn)
}

func numberOfSteps(num int) int {
	count := 0
	for {
		if num == 0 {
			break
		} else if num%2 == 0 {
			num /= 2
			count += 1
		} else {
			num -= 1
			count += 1
		}
	}
	return count
}

func NameGenerator() string {
	rand.Seed(time.Now().UnixNano())
	seed := time.Now().UnixMicro()
	name := namegenerator.NewNameGenerator(seed)

	return name.Generate()
}
