// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package server

import (
	"context"
	logic2 "douyinbackend/app/usercenter/cmd/rpc/internal/logic"
	"douyinbackend/app/usercenter/cmd/rpc/internal/svc"
	pb2 "douyinbackend/app/usercenter/cmd/rpc/pb/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb2.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) RegisterUser(ctx context.Context, in *pb2.DouyinUserRegisterRequest) (*pb2.DouyinUserRegisterResponse, error) {
	l := logic2.NewRegisterUserLogic(ctx, s.svcCtx)
	return l.RegisterUser(in)
}

func (s *UsercenterServer) LoginUser(ctx context.Context, in *pb2.DouyinUserLoginRequest) (*pb2.DouyinUserLoginResponse, error) {
	l := logic2.NewLoginUserLogic(ctx, s.svcCtx)
	return l.LoginUser(in)
}

func (s *UsercenterServer) GetUser(ctx context.Context, in *pb2.DouyinUserRequest) (*pb2.DouyinUserResponse, error) {
	l := logic2.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}
