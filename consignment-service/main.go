package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-demo/consignment-service/proto/consignment"
	"net"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// 模拟数据库
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// service中要实现pb中所有的方法
type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// 保存我们的consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &Repository{}

	// 增加tcp的端口监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	// 创建一个服务
	s := grpc.NewServer()
	// 将service注册到服务中
	pb.RegisterShippingServiceServer(s, &service{repo: repo})
	// 服务注册reflection
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
