package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/greet/internal/config"
	"go-zero-demo/greet/internal/handler"
	"go-zero-demo/greet/internal/svc"
)

func main() {
	var configFile = flag.String("f", "greet/etc/greet-api.yaml", "the config file")

	var c config.Config
	conf.MustLoad(*configFile, &c)
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}

func swap(a, b *int) {
	fmt.Printf("a=%d,b=%d\n", a, b)
	fmt.Printf("a=%d,b=%d\n", *a, *b)
	*a, *b = *b, *a
}
