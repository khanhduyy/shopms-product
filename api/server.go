package api

import (
	"context"

	"github.com/khanhduyy/shopms-product/config"
	"github.com/khanhduyy/shopms-product/internal/repository"
	"github.com/khanhduyy/shopms-product/internal/service"
	pb "github.com/khanhduyy/shopms-product/rpc/product"

	"github.com/khanhduyy/shopms-common/client/db"
	"github.com/khanhduyy/shopms-common/grpc"
	log "github.com/khanhduyy/shopms-common/logger"
)

var (
	ctx = context.Background()
)

func Run() {
	config.Load()

	log.New(&log.Config{
		Level:  config.Values.LoggerLevel,
		Format: "json",
		Name:   "shopms-product",
	})

	log.Infof("Start running product service")

	dbConfig := &db.Config{
		Host:         config.Values.DB.Host,
		Port:         int(config.Values.DB.Port),
		DatabaseName: config.Values.DB.DBName,
		User:         config.Values.DB.User,
		Password:     config.Values.DB.Password,
	}
	clientDB, err := db.Open(*dbConfig)
	if err != nil {
		log.Panicf("db open connection on failure ", err)
	}

	defer func() {
		if err := clientDB.Close(); err != nil {
			panic(err)
		}
	}()

	productApi := initProductApi(clientDB)

	grpcService := grpc.StandardRunner{
		Address: config.Values.ServerAddress,
		Server:  grpc.New(ctx),
	}

	pb.RegisterProductAPIServer(grpcService.Server, productApi)

	if err := grpcService.Run(ctx); err != nil {
		log.Panicf("failed to serve the schedule service", err)
	}
}

func initProductApi(client *db.Client) *ProductApi {
	productRepository := repository.NewProductRepository(client)
	productService := service.NewProductService(productRepository)
	productEndpoint := &ProductApi{
		ProductService: productService,
	}
	return productEndpoint
}
