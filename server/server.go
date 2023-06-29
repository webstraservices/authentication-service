package server

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
	resp := pb.ValidateResponse{
		Status: http.StatusUnauthorized,
		Error:  "failed to validate token",
		UserId: 0,
	}
	if req.Token == "fail" {
		return &resp, nil
	}

	resp.UserId = 1

	return &resp, nil
}
