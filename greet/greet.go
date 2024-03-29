package main

import (
	"flag"
	"fmt"

	"learnGoWithTest/learn_go_with_tests/greet/internal/config"
	"learnGoWithTest/learn_go_with_tests/greet/internal/handler"
	"learnGoWithTest/learn_go_with_tests/greet/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
