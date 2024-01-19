package food

import (
	"context"
	"fmt"

	pb "grpc-gorm-mysql/proto/food"
)

type MyServer struct {
	pb.UnimplementedOperationServer
	FoodRepo *FoodRepository
}

func (m *MyServer) Insert(ctx context.Context, in *pb.InsDelUpdRequest) (*pb.Reply, error) {
	food := Food{
		Id:         in.GetId(),
		Name:       in.GetName(),
		Price:      in.GetPrice(),
		TypeId:     in.GetTypeId(),
		CreateTime: in.GetCreateTime(),
	}

	// Chame o método Insert no repositório de alimentos
	m.FoodRepo.Insert(food)
	return &pb.Reply{Result: "Insert completed."}, nil
}

func (m *MyServer) Delete(ctx context.Context, in *pb.InsDelUpdRequest) (*pb.Reply, error) {
	food := Food{
		Id:         in.GetId(),
		Name:       in.GetName(),
		Price:      in.GetPrice(),
		TypeId:     in.GetTypeId(),
		CreateTime: in.GetCreateTime(),
	}

	// Chame o método Insert no repositório de alimentos
	m.FoodRepo.Delete(food)
	return &pb.Reply{Result: "Delete Completed"}, nil
}
func (m *MyServer) Update(ctx context.Context, in *pb.InsDelUpdRequest) (*pb.Reply, error) {
	food := Food{
		Id:         in.GetId(),
		Name:       in.GetName(),
		Price:      in.GetPrice(),
		TypeId:     in.GetTypeId(),
		CreateTime: in.GetCreateTime(),
	}
	m.FoodRepo.Update(food)
	return &pb.Reply{Result: "Update completed."}, nil
}
func (m *MyServer) Select(ctx context.Context, in *pb.SelectRequest) (*pb.Reply, error) {
	result := m.FoodRepo.Select(in.GetColumns(), in.GetCondition())
	fmt.Println(result)
	return &pb.Reply{Result: "Select completed."}, nil
}

func (m *MyServer) ExecSql(ctx context.Context, in *pb.SqlRequest) (*pb.Reply, error) {
	m.FoodRepo.ExecSql(in.GetSql())
	return &pb.Reply{Result: "Execution completed."}, nil
}
