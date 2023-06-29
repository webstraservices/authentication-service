package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/webstrasuite/webstra-auth/pb"
	"github.com/webstrasuite/webstra-auth/server"
	"google.golang.org/grpc"
)

const defaultPort = ":3001"

func main() {
	// Load .env file if necessary
	if os.Getenv("DB_CONNECTION_STRING") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Unable to load environment files - Err: %s", err)
		}
	}

	ln, err := net.Listen("tcp", defaultPort)
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}

	fmt.Println("Webstra Auth service running on", defaultPort)

	s := server.Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(ln); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
