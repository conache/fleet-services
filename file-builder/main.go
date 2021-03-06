package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/Condition17/fleet-services/file-builder/chunks-storage"
	"github.com/Condition17/fleet-services/file-builder/config"
	"github.com/Condition17/fleet-services/file-builder/handler"
	proto "github.com/Condition17/fleet-services/file-builder/proto/file-builder"
	"google.golang.org/grpc"
	"log"
	"net"
)

const ChunksBucketName = "fleet-files-chunks"

func main() {
	var lis net.Listener
	var conn *grpc.ClientConn
	var pubSubClient *pubsub.Client
	var err error
	configs := config.GetConfig()

	// Server startup
	if lis, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", configs.ServerPort)); err != nil {
		log.Fatalf("Server failed to listen on port ':%v'. Error encountered: %v\n", configs.ServerPort, err)
	}
	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	// Create connection to Fleet micro services proxy
	if conn, err = grpc.Dial(configs.FleetServicesGrpcProxyUrl, grpc.WithInsecure()); err != nil {
		log.Fatalf("Error encountered while creating connection to fleet services grpc proxy: %v", err)
	}
	log.Println("Connection to fleet services GRPC proxy successfully initiated")
	defer conn.Close()

	// Create Pub Sub client
	pubSubClient, err = pubsub.NewClient(context.Background(), configs.GoogleProjectID)
	if err != nil {
		log.Fatalf("Error creating Google Pub Sub client: %v", err)
	}

	// Create chunks storage client
	chunksStorageClient, err := chunksStorage.NewChunksStorageClient(configs.GoogleProjectID, ChunksBucketName)
	if err != nil {
		log.Fatalf("Error encountered while initializing the chunks storage client: %v", err)
	}
	proto.RegisterFileBuilderServer(grpcServer, handler.NewHandler(conn, *pubSubClient, *chunksStorageClient))

	log.Printf("Starting GRPC server on localhost:%v\n", configs.ServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
