package main

import (
	"context"
	"fmt"
	"net"
	"protoserver/proto"

	"github.com/fxtlabs/primes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type MathServer struct{}

func Fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibo(n-2) + Fibo(n-1)
}

func (s *MathServer) Fibo(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a := request.GetA()
	n := int(a)
	result := ""
	for i := 0; i < n; i++ {
		if result == "" {
			result += fmt.Sprint(Fibo(i))
		} else {
			result += " " + fmt.Sprint(Fibo(i))
		}
	}
	return &proto.Response{Result: result}, nil
}
func (s *MathServer) IsPrime(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a := request.GetA()
	result := fmt.Sprintf("%t", primes.IsPrime(int(a)))
	return &proto.Response{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":6060")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterMathServiceServer(srv, &MathServer{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
