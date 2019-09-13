package server

import (
	"context"
	"database/sql"

	// "flag"
	"fmt"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/TranThienChi/go-grpc-http-rest-microservice/pkg/protocol/grpc"
	"github.com/TranThienChi/go-grpc-http-rest-microservice/pkg/protocol/rest"
	v1 "github.com/TranThienChi/go-grpc-http-rest-microservice/pkg/service/v1"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	// flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	// flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	// flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	// flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	// flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	// flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	// flag.Parse()

	// if len(cfg.GRPCPort) == 0 {
	cfg.GRPCPort = os.Getenv("GRPCPort")
	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}
	// }

	// if len(cfg.HTTPPort) == 0 {
	cfg.HTTPPort = os.Getenv("HTTPPort")
	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}
	// }

	// if len(cfg.DatastoreDBHost) == 0 {
	cfg.DatastoreDBHost = os.Getenv("DatastoreDBHost")
	// }

	// if len(cfg.DatastoreDBUser) == 0 {
	cfg.DatastoreDBUser = os.Getenv("DatastoreDBUser")
	// }

	// if len(cfg.DatastoreDBPassword) == 0 {
	cfg.DatastoreDBPassword = os.Getenv("DatastoreDBPassword")
	// }

	// if len(cfg.DatastoreDBSchema) == 0 {
	cfg.DatastoreDBSchema = os.Getenv("DatastoreDBSchema")
	// }

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	v1API := v1.NewToDoServiceServer(db)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
