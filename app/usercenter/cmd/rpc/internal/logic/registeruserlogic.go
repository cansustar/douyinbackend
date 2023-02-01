package logic

import (
	"context"
	"douyinbackend/app/usercenter/cmd/rpc/internal/svc"
	"douyinbackend/app/usercenter/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterUserLogic) RegisterUser(in *pb.DouyinUserRegisterRequest) (*pb.DouyinUserRegisterResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DouyinUserRegisterResponse{}, nil
}
