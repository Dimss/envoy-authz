package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func startServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", viper.GetString("api.grpc.address"), viper.GetString("api.grpc.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Infof("grpc server listening on %s:%s", viper.GetString("api.grpc.address"), viper.GetString("api.grpc.port"))
	//tlsCredentials, err := loadServerTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	opts := []grpc.ServerOption{
		//grpc.UnaryInterceptor(unaryServerInterceptor()),
		//grpc.StreamInterceptor(streamServerInterceptor()),
		//grpc.Creds(tlsCredentials),
	}
	grpcServer := grpc.NewServer(opts...)
	//registerGrpcServices(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
