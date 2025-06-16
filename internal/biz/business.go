package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)


type BusinessRepo interface {
	Save(context.Context) error
}

type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

