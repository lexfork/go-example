package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gunsluo/go-example/grpc/ssl/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	sAddress = "0.0.0.0:3264"
	crtFile  = "../cert/server.crt"
	keyFile  = "../cert/server.key"
)

type Service struct {
}

func (s *Service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "no metadata")
	}

	token := md.Get("token")
	if len(token) == 0 {
		return nil, grpc.Errorf(codes.Unauthenticated, "no token")
	}

	fmt.Println("requst:", token[0], req.Name)
	return &pb.HelloReply{
		Message: "hello, " + req.Name,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", sAddress)
	if err != nil {
		panic(err)
	}

	// Create the TLS credentials
	var opts []grpc.ServerOption
	if crtFile != "" && keyFile != "" {
		fmt.Println("enable credentials in the grpc")
		creds, err := credentials.NewServerTLSFromFile(crtFile, keyFile)
		if err != nil {
			panic(err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	server := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(server, &Service{})

	logrus.WithField("addr", sAddress).Println("Starting server")
	server.Serve(listener)
}
