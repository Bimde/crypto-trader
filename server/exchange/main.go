package main

import (
	"context"
	"log"
	"net"
	"fmt"

	"github.com/Bimde/crypto-trader/build"
	"github.com/Bimde/httputils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var http httputils.Client

const api = "https://api.cryptonator.com/api/ticker/%s"

func main() {
	lis, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting stock exchange")
	s := grpc.NewServer()
	build.RegisterStockExchangeServer(s, &server{})
	reflection.Register(s)
	http = httputils.Client{}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) GetStockPrice(c context.Context, req *build.GetStockPriceRequest) (*build.GetStockPriceResponse, error) {
	var response build.CryptonatorResponse
	http.Get(nil, getAPIURL(req.Ticker[0]), nil, response)
	if response.Ticker != nil {
		return &build.GetStockPriceResponse{Stock: []*build.Stock{&build.Stock{Ticker: response.Ticker.Base, Price: response.Ticker.Price}}}, nil
	}
	return nil, fmt.Errorf("Error getting quote for ticker %s", req.Ticker[0])
}

func getAPIURL(ticker string) string {
	return fmt.Sprintf(api, ticker)
}
