package main

import (
	"context"
	"errors"
	"fmt"
	pb "go_grpc_demo/pkg/agenda_server/v1"
	"go_grpc_demo/pkg/service"
	"go_grpc_demo/pkg/setupotel"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := 8082
	lis, err := net.Listen("tcp", fmt.Sprintf("[::]:%d", port))

	if err != nil {
		panic(err)
	}

	// Handling graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()

	// Set up OpenTelemetry.
	serviceName := "agenda"
	serviceVersion := "0.1.0"
	otelShutdown, err := setupotel.SetOTelSDK(ctx, serviceName, serviceVersion)

	if err != nil {
		log.Fatalf("failed to set up OpenTelemetry: %v", err)
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	grpcServer := grpc.NewServer()

	srv, err := service.NewService()

	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	pb.RegisterAgendaServiceServer(grpcServer, srv)

	// Create a graceful shutdown
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Printf("Server started on port %d", port)

	<-ch // Block until a signal is received

	// Wait for the interrupt signal to gracefully shut down the server with a timeout of 5 seconds.
	grpcServer.GracefulStop()
	log.Println("Shutting down server...")
	time.Sleep(5 * time.Second)
	defer func() {
		err = srv.Close()
		if err != nil {
			log.Panicf("error shutting down service: %v\n", err)
		}
	}()
	log.Println("Server stopped")

}
