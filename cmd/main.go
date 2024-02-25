package main

import (
	server "assignment-permission/cmd/server"
	api2 "assignment-permission/internal/api"
	"assignment-permission/internal/config"
	"assignment-permission/internal/permission"
	mongo "assignment-permission/internal/pkg"
	"context"
	"fmt"
)

var exitErr error

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig(".", "app_config")
	if err != nil {
		fmt.Println("Error loading config: %+v", err)
		panic(err)
	}

	fmt.Println("Starting server with config: %+v", cfg)
	// init mongo client to be used by the service
	mongoCfg := mongo.Config{
		ConnectionURI: cfg.DatabaseURI,
		DatabaseName:  "assignment-permission",
	}
	mi, err := initMongoClient(ctx, mongoCfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Mongo client connected")

	fmt.Println("Initializing service")
	// init service to be used by the api
	service := initService(ctx, mi, permission.MongoConfig{
		PermissionCollection: cfg.MongoConfig.PermissionCollection,
		RoleCollection:       cfg.MongoConfig.RoleCollection,
	})
	fmt.Println("Service initialized")

	// init api to be used by the http server
	api := api2.New(*service)
	fmt.Println("Configuration loaded: %+v", cfg)
	httpServer := server.New(&server.Config{
		Host: cfg.Host,
		Port: cfg.Port,
	}, api)
	fmt.Println("Starting HTTP server on port %d", cfg.Port)
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
