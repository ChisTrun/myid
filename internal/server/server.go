package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb0 "github.com/ChisTrun/myid/api"
	"github.com/ChisTrun/myid/internal/server/myid"
	"github.com/ChisTrun/myid/package/config"
	"github.com/ChisTrun/myid/package/ent"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Config) {
	fmt.Println(cfg.Database)
	client, err := ent.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb0.RegisterMyIdServer(grpcServer, myid.NewServer())
	log.Printf("server is runing on: %v:%v", cfg.Server.Host, cfg.Server.Port)
	grpcServer.Serve(lis)
}
