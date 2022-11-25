package service

import "go-zero-demo/product/rpc/svc"

type ProductService struct {
	svcCtx *svc.ServiceContext
}

func NewProductService(ctx *svc.ServiceContext) *ProductService {
	return &ProductService{
		svcCtx: ctx,
	}
}
