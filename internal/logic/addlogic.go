package logic

import (
	"context"

	"github.com/jarcn/bookadd/add"
	"github.com/jarcn/bookadd/internal/svc"
	"github.com/jarcn/bookadd/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	// 手动代码开始
    _, err := l.svcCtx.Model.Insert(l.ctx,&model.Book{
        Book:  in.Book,
        Price: in.Price,
    })
    if err != nil {
        return nil, err
    }

    return &add.AddResp{
        Ok: true,
    }, nil
    // 手动代码结束
}
