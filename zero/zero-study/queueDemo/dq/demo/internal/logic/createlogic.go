package logic

import (
	"context"
	"time"

	"binggo/zero/zero-study/queueDemo/dq/demo/internal/svc"
	"binggo/zero/zero-study/queueDemo/dq/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	msg := "unPay, cancel order"

	// 5s后推送消息
	deplayResp, err := l.svcCtx.DqPusherClient.Delay([]byte(msg), time.Second*5)
	if err != nil {
		logx.Errorf("error from DqPusherClient Delay err : %v", err)
	}
	logx.Infof("resp : %s", deplayResp) // fmt.Sprintf("%s/%s/%d", p.endpoint, p.tube, id)

	logx.Info("没有阻塞！")
	//// 2、在某个指定时间执行
	//atResp, err := l.svcCtx.DqPusherClient.At([]byte(msg), time.Now())
	//if err != nil {
	//	logx.Errorf("error from DqPusherClient Delay err : %v", err)
	//}
	//logx.Infof("resp : %s", atResp) // fmt.Sprintf("%s/%s/%d", p.endpoint, p.tube, id)

	return
}
