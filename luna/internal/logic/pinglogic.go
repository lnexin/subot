package logic

import (
	"context"

	"su-wcbot/luna/internal/svc"
	"su-wcbot/luna/wcbot"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *wcbot.Request) (*wcbot.Response, error) {
	// todo: add your logic here and delete this line

	//return &luna.Response{}, nil
	return &wcbot.Response{Pong: "test"}, nil
}
