package logic

import (
	"context"
	"douyinbackend/app/usercenter/cmd/rpc/internal/svc"
	"douyinbackend/app/usercenter/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginUserLogic) LoginUser(in *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DouyinUserLoginResponse{}, nil
}
