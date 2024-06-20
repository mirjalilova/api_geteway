package main

import (
	"api-geteway/api"
	cf "api-geteway/config"
	"fmt"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := cf.Load()

	Conn, err := grpc.NewClient(fmt.Sprintf("localhost%s", config.LIBRARY_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Error while connecting to reservation service: " + err.Error())
		return
	}
	defer Conn.Close()

	r := api.NewRouter(Conn)

	slog.Info("Starting HTTP server on port " + config.HTTP_PORT + "...")
	if r.Run(config.HTTP_PORT); err != nil {
		slog.Error("Error while starting HTTP server: " + err.Error())
	}
}
