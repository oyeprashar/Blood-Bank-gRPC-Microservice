package main

import (
	proto "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/setup"
	"code.nurture.farm/BloodBankSystemService/core/golang/hook"
	"flag"
	"fmt"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	
)

func runMonitoring(grpcServer *grpc.Server) {
	// register prometheus
	grpcPrometheus.Register(grpcServer)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", "0.0.0.0", 6005), nil)
	if err != nil {
		logger.Panic("Unable to start prometheus handler", zap.Error(err))
	}
}



func main() {

	port := flag.Int("port", 6000, "Port for GRPC server to listen")
	flag.Parse()
	logger.Info("Starting Farm Service Service!")
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		logger.Fatal("Unable to listen on port", zap.Int("port", *port), zap.Error(err))
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcPrometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpcPrometheus.UnaryServerInterceptor))
	proto.RegisterBloodBankSystemServiceServer(grpcServer, setup.BloodBankSystemService)
	logger.Info("Registered server",
		zap.Any("grpcServer", grpcServer), zap.Any("listener", lis), zap.Int("port", *port))

	// on GRPC services
	go runMonitoring(grpcServer)

	

	hook.PreStartUpHook()

	// Start server
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal("Unable to listen on service", zap.Int("port", *port), zap.Error(err))
	}

	hook.PostStartUpHook()
}
