package main

import (
	"github.com/preethi29/go-grpc/order"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) PlaceOrder(ctx context.Context, orderRequest *order.OrderRequest) (*order.OrderResponse, error) {
	if orderRequest.ItemId != 0 {
		return &order.OrderResponse{OrderPlaced: true}, nil
	}
	return &order.OrderResponse{OrderPlaced: false}, nil

}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	order.RegisterOrderServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
