package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao ouvir na porta: %v", err)
	}

	s := grpc.NewServer()

	fmt.Println("Servidor gRPC iniciado na porta :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao servir: %v", err)
	}
}
