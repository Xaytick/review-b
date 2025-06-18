package data

import (
	"context"

	v1 "review-b/api/review/v1"
	"review-b/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}


// NewBusinessRepo .
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *businessRepo) Reply(ctx context.Context, param *biz.ReplyParam) (int64, error) {
	r.log.WithContext(ctx).Infof("[data] Reply: %+v", param)
	// 之前是直接操作数据库, 现在需要调用rpc接口来实现
	ret, err :=r.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewID: param.ReviewID,
		StoreID: param.StoreID,
		Content: param.Content,
		PicInfo: param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	r.log.WithContext(ctx).Debugf("[data] ReplyReview: ret:%+v, err:%+v", ret, err)
	if err != nil {
		return 0, err
	}
	return ret.ReplyID, nil
}

func (r *businessRepo) AppealReview(ctx context.Context, param *biz.AppealReviewParam) (int64, error) {
	r.log.WithContext(ctx).Infof("[data] AppealReview: %+v", param)
	ret, err := r.data.rc.AppealReview(ctx, &v1.AppealReviewRequest{
		ReviewID: param.ReviewID,
		StoreID: param.StoreID,
		Reason: param.Reason,
		Content: param.Content,
		PicInfo: param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	r.log.WithContext(ctx).Debugf("[data] AppealReview: ret:%+v, err:%+v", ret, err)
	if err != nil {
		return 0, err
	}
	return ret.AppealID, nil
}
