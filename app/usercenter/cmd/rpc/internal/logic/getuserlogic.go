package logic

import (
	"context"
	"douyinbackend/app/usercenter/cmd/rpc/internal/svc"
	"douyinbackend/app/usercenter/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.DouyinUserRequest) (*pb.DouyinUserResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DouyinUserResponse{}, nil
}
