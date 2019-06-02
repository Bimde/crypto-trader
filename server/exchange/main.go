package main

import (
	"context"
	"log"
	"net"

	"../../proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const ()

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting Cart service")
	s := grpc.NewServer()
	pb.RegisterStockExchangeServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func GetStockPrice(c context.Context, req *GetStockPriceRequest) (*GetStockPriceResponse, error)
	return &GetStockPriceResponse{Stock{Ticker: "GOOG", Price: 21.22}}, nil
}
