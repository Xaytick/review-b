package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)


type BusinessRepo interface {
	Reply(context.Context, *ReplyParam) (int64, error)
	AppealReview(context.Context, *AppealReviewParam) (int64, error)
}

type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

// 创建回复, 在service中调用
func (uc *BusinessUsecase) CreateReply(ctx context.Context, param *ReplyParam) (int64, error) {
	uc.log.WithContext(ctx).Infof("[biz] CreateReply: %+v", param)
	return uc.repo.Reply(ctx, param)
}

// 申诉评论, 在service中调用
func (uc *BusinessUsecase) AppealReview(ctx context.Context, param *AppealReviewParam) (int64, error) {
	uc.log.WithContext(ctx).Infof("[biz] AppealReview: %+v", param)
	return uc.repo.AppealReview(ctx, param)
}




