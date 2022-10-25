package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/ninja-dark/fibonacci_testtask/config"
	"github.com/ninja-dark/fibonacci_testtask/internal/fiboLogic"
	grpcsrv "github.com/ninja-dark/fibonacci_testtask/internal/infrastructure/grpcSrv"
	"github.com/ninja-dark/fibonacci_testtask/internal/infrastructure/rest/api/handler"
	"github.com/ninja-dark/fibonacci_testtask/internal/infrastructure/rest/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)
	const(
		configPath = "config/configs.yaml"
	)

func main() {
	//set logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	path := flag.String("config", configPath, "file path")
	flag.Parse()
	//parse config
	cfg, err := config.ParseConfig(*path)
	if err != nil {
		logrus.Fatalf("error initializating configs: %s" , err.Error())
	}
	restApi := cfg.Rest
	grpcPort := cfg.GrpcPort
	cache := cfg.Memcache
	
	m := memcache.New(cache)
	if err := m.Ping(); err != nil {
		logrus.Error("Cannot to connect to memcache srv")
	}
	
	f:= fibologic.Fibo{Cache: m}
	//srart rest api
	srv := new(server.Server)
	rest := handler.Handler{
		Services: &f,
	}
	go func() {
		if err := srv.Run(restApi, rest.InitRouters()); err != nil{
			logrus.Fatal("error occured while running rest api server: %s", err.Error())
		}
	}()
	logrus.Print("Fibonacci Started")
	// start grps
	grpcNew:= grpc.NewServer()
	grpcSrv:= &grpcsrv.ServerG{Handler: &f}

	go grpcsrv.Run(grpcNew, grpcSrv, grpcPort)
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Print("Fibonacci Shitting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

