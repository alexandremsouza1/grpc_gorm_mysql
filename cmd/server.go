package main

import (
	"fmt"
	"net"

	"grpc-gorm-mysql/app/food"
	"grpc-gorm-mysql/mysql"
	pb "grpc-gorm-mysql/proto/food"

	"google.golang.org/grpc"
)

const port = ":8080"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}

	server := grpc.NewServer()
	mysql.Connect()
	db := mysql.GetDB()
	repository := food.NewFoodRepository(db)
	pb.RegisterOperationServer(server, &food.MyServer{FoodRepo: repository})

	fmt.Println("grpc service starts...")
	server.Serve(list)
}
