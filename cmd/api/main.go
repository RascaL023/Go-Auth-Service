package main

import (
	"context"
	"fmt"
	"go-auth-service/internal/config"
	"go-auth-service/internal/domain/authority"
	"go-auth-service/internal/infra/database"
	appHtpp "go-auth-service/internal/transport/http/authority"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	sigCh				:= make(chan os.Signal, 1) // Buffer 1 untuk terima 1 input terminal dengan tipe os.Signal
	signal.Notify(sigCh, os.Interrupt) // Kirim sinyal interupt e.g ctrl + c(SIGINT) ke sigCh
	parentCtx 			:= context.Background()

	startupCtx, startupCancel := context.WithTimeout(
		parentCtx,
		3*time.Second,
	)
	defer startupCancel()

	config, err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	authorityRepo, err := database.ResolveAuthorityRepository(startupCtx, config)
	if  err != nil {
		log.Fatal(err.Error())
		return
	}

	authorityService	:= authority.NewService(authorityRepo)	
	authorityHandler	:= appHtpp.NewAuthorityHandler(authorityService)

	mux 				:= http.NewServeMux()
	authorityRouter		:= appHtpp.NewAuthorityRouter(mux, authorityHandler)

	authorityRouter.Handle()

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", config.Server.Port),
		Handler: mux,
	}
	go server.ListenAndServe();
	fmt.Println("App run on ", config.Server.Port)

	<-sigCh
	shutdownCtx, shutdownCancel	:= context.WithTimeout(
		parentCtx,
		5*time.Second,
	)
	defer shutdownCancel()

	server.Shutdown(shutdownCtx)
	fmt.Println("App stoped")
}
