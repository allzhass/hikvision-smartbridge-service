package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"smartbridge-service"
	"smartbridge-service/pkg/handler"
	"smartbridge-service/pkg/service"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	logrus.Println(os.Getenv("SERVER_PORT"))
	logrus.Println(os.Getenv("SMART_BRIDGE_URL"))

	services := service.NewService(os.Getenv("SMART_BRIDGE_URL"))
	handlers := handler.NewHandler(services)

	httpsServer := smartbridge_service.HttpsServer{
		Port: os.Getenv("SERVER_PORT"),
	}
	srv := new(smartbridge_service.Server)
	go func() {
		if err := srv.Run(httpsServer, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Todo Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Todo Server Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
