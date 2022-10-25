package grpcsrv

import (
	"context"
	"fmt"
	"net"

	fibologic "github.com/ninja-dark/fibonacci_testtask/internal/fiboLogic"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)
type ServerG struct{
	Handler *fibologic.Fibo
	UnimplementedFibonacciServer
}
func (s *ServerG)FiboGRPC(ctx context.Context, in *FiboRequest) (*FiboResponse, error){
	r, err := s.Handler.GetSequence(int(in.GetX()), int(in.GetY()))
	if err != nil {
		return nil, err
	}
	return &FiboResponse{Result: r}, nil
}

func Run(srv *grpc.Server, srvFib FibonacciServer, port string){
	RegisterFibonacciServer(srv, srvFib)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil{
		logrus.Fatalf("failed to listen: %v", err)
	}
	logrus.Print("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil{
		logrus.Fatal("grpc failed to serve: %v", err)
	}
}