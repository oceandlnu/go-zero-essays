package logic

import (
	"context"

	"go_zero_essays/internal/svc"
	"go_zero_essays/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Go_zero_essaysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGo_zero_essaysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Go_zero_essaysLogic {
	return &Go_zero_essaysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Go_zero_essaysLogic) Go_zero_essays(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{}
	resp.Message = req.Name
	return
}
