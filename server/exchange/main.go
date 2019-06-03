package main

import (
	"context"
	"log"
	"net"

	"github.com/Bimde/crypto-trader/build"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting stock exchange")
	s := grpc.NewServer()
	build.RegisterStockExchangeServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) GetStockPrice(c context.Context, req *build.GetStockPriceRequest) (*build.GetStockPriceResponse, error) {
	return &build.GetStockPriceResponse{Stock: []*build.Stock{&build.Stock{Ticker: "GOOG", Price: 21.22}}}, nil
}
