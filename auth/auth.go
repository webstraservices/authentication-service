package auth

import (
	"context"
	"net/http"

	"github.com/webstrasuite/webstra-auth/pb"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  "tempTokenString",
	}, nil
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: 1,
	}, nil
}
