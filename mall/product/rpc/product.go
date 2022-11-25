package main

import (
	"github.com/liudng/godump"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/product/rpc/config"
	"go-zero-demo/product/rpc/service"
	"go-zero-demo/product/rpc/svc"
	"google.golang.org/grpc"
)

func main() {
	configFile := "product/rpc/etc/product.yaml"
	var c config.Config
	conf.MustLoad(configFile, &c)

	ctx := svc.NewServiceContext(c)

	zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		grpcServer.RegisterService(&User_ServiceDesc, service.NewProductService(ctx))
	})
	godump.Dump(c)
}
