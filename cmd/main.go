package main

import (
	server "assignment-permission/cmd/server"
	api2 "assignment-permission/internal/api"
	"assignment-permission/internal/config"
	"assignment-permission/internal/permission"
	mongo "assignment-permission/internal/pkg"
	"context"
)

var exitErr error

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig(".", "app_config")
	if err != nil {
		panic(err)
	}

	// init mongo client to be used by the service
	mongoCfg := mongo.Config{
		ConnectionURI: cfg.DatabaseURI,
		DatabaseName:  "assignment-permission",
	}
	mi, err := initMongoClient(ctx, mongoCfg)
	if err != nil {
		panic(err)
	}

	// init service to be used by the api
	service := initService(ctx, mi, permission.MongoConfig{
		PermissionCollection: "permissions",
	})

	// init api to be used by the http server
	api := api2.New(*service)

	httpServer := server.New(&server.Config{
		Host: cfg.Host,
		Port: cfg.Port,
	}, api)

	fatalErr := make(chan error, 1)
	startServices(fatalErr, httpServer)

	exitErr = <-fatalErr
}

func initMongoClient(ctx context.Context, cfg mongo.Config) (mi *mongo.MongoInstance, err error) {
	mi, err = mongo.NewMongoClient(cfg)
	return mi, err
}

func initService(ctx context.Context, mi *mongo.MongoInstance, cfg permission.MongoConfig) (service *permission.Service) {
	persistence, err := permission.InitDB(mi, &cfg)
	if err != nil {
		panic(err)
	}
	newService, err := permission.NewService(persistence)
	if err != nil {
		return nil
	}
	return &newService
}

func startServices(fatalErr chan error, httpServer *server.HTTP) {
	go func() {
		fatalErr <- httpServer.Start()
	}()
}
