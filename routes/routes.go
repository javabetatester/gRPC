package routes

import (
	"google.golang.org/grpc"
	"grpc-app/handlers"
)

func RegisterRoutes(s *grpc.Server, handler *handlers.GRPCHandler) {
	// pb.RegisterUserServiceServer(s, handler.UserController)
}
