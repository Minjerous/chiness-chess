package service

import (
	"chess-user/cmd/rpc/internal/logic"
	"chess-user/cmd/rpc/proto/user"
	"context"
)

type UserServer struct {
	user.UnimplementedUserServiceServer
}

func (s *UserServer) UserLogin(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	return logic.Login(req)
}

func (s *UserServer) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	return logic.Register(req)
}

func (s *UserServer) UserInfo(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {

	panic("Unimplemented!")
}
