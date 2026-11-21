package main

import (
	"context"
	"errors"
	"fmt"
	"go-auth-service/internal/config"
	"go-auth-service/internal/domain/authority"
	"go-auth-service/internal/infra/database"
	authorityhttp "go-auth-service/internal/transport/http/authority"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	sigCh := make(chan os.Signal, 1)   // Buffer 1 untuk terima 1 input terminal dengan tipe os.Signal
	signal.Notify(sigCh, os.Interrupt) // Kirim sinyal interupt e.g ctrl + c(SIGINT) ke sigCh
	defer signal.Stop(sigCh)
	parentCtx := context.Background()

	startupCtx, startupCancel := context.WithTimeout(
		parentCtx,
		3*time.Second,
	)
	defer startupCancel()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	authorityRepo, err := database.ResolveAuthorityRepository(startupCtx, cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	authorityService := authority.NewService(authorityRepo)
	authorityHandler := authorityhttp.NewAuthorityHandler(authorityService)

	mux := http.NewServeMux()
	authorityRouter := authorityhttp.NewAuthorityRouter(mux, authorityHandler)

	authorityRouter.Handle()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: mux,
	}

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", server.Addr, err)
	}

	serveErrCh := make(chan error, 1)
	go func() {
		serveErrCh <- server.Serve(listener)
	}()
	log.Printf("server listening on %s", listener.Addr())

	select {
	case err := <-serveErrCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	case <-sigCh:
		shutdownCtx, shutdownCancel := context.WithTimeout(
			parentCtx,
			5*time.Second,
		)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
			_ = server.Close()
		}

		log.Println("server stopped")
	}
}
