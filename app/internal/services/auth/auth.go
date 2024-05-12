package auth

import (
	"context"
	pb "github.com/k-vanio/gRPC-with-docker/internal/pb/auth"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(ctx context.Context, in *pb.Credentials) (*pb.Response, error) {
	return &pb.Response{Token: "token"}, nil
}
