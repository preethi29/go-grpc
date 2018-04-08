package main

import (
	"github.com/preethi29/go-grpc/order"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := order.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.PlaceOrder(ctx, &order.OrderRequest{ItemId: 1, ItemName: "Shampoo", ItemQuantity: 3})
	if err != nil {
		log.Fatalf("could not place order: %v", err)
	}
	log.Printf("Order status: %s", resp.OrderPlaced)
}
