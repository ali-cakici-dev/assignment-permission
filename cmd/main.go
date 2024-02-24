package main

import (
	"assignment-permission/internal/config"
	mongo "assignment-permission/internal/pkg"
	"context"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig(".", "app_config")
	if err != nil {
		panic(err)
	}
	mongoCfg := mongo.Config{
		ConnectionURI: cfg.DatabaseURI,
		DatabaseName:  "assignment-permission",
	}

	mi, err := initMongoClient(ctx, mongoCfg)
	if err != nil {
		panic(err)
	}

	service := initService(ctx, mi)

}

func initMongoClient(ctx context.Context, cfg mongo.Config) (mi *mongo.MongoInstance, err error) {
	mi, err = mongo.NewMongoClient(cfg)
	return mi, err
}

func initService(ctx context.Context, mi *mongo.MongoInstance) (service *assignment_permission.Service) {
	service = &assignment_permission.Service{
		Mongo: mi,
	}
	return service
}
