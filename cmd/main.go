package main

import (
	"fmt"
	"log"

	"github.com/juliofilizzola/Code-pix/application/grpc"
	"github.com/juliofilizzola/Code-pix/config/env"
	"github.com/juliofilizzola/Code-pix/infra/db"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func main() {
	env.LoadingEnv()
	fmt.Print(env.Dns)
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("server db not initialized")
	}
	grpc.StartGrpcServer(database, 50051)

}
