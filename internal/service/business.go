package service

import (
	"context"
	"fmt"

	pb "review-b/api/business/v1"
	"review-b/internal/biz"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer
	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

func (s *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	replyId, err := s.uc.CreateReply(ctx, &biz.ReplyParam{
		ReviewID: req.ReviewID,
		StoreID: req.StoreID,
		Content: req.Content,
		PicInfo: req.PicInfo,
		VideoInfo: req.VideoInfo,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ReplyReviewReply{ReplyID: replyId}, nil
} 

// AppealReview 申诉评论
func (s *BusinessService) AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error) {
	fmt.Println("[service] AppealReview, req:", req)
	// 调用biz层
	appealId, err := s.uc.AppealReview(ctx, &biz.AppealReviewParam{
		ReviewID:  req.ReviewID,
		StoreID:   req.StoreID,
		Reason:    req.Reason,
		Content:   req.Content,
		PicInfo:   req.PicInfo,
		VideoInfo: req.VideoInfo,
	})
	if err != nil {
		return nil, err
	}
	// 拼装返回值
	return &pb.AppealReviewReply{AppealID: appealId, Status: 10}, nil
}