package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Jidetireni/tiny/config"
	"github.com/Jidetireni/tiny/internals/shorten"
	"github.com/Jidetireni/tiny/pkg/cassandra"
	"github.com/Jidetireni/tiny/pkg/zookeeper"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// 1. Initialize dependencies (The "Ask for it" philosophy starts here)
	config := config.New()
	zookeeper, err := zookeeper.New(config)
	if err != nil {
	}

	cassandra := cassandra.New()

	shortenService := shorten.New(zookeeper, cassandra)

	srv := NewServer(
		config,
		shortenService,
	)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.ServerConfig.Host, config.ServerConfig.Port),
		Handler: srv,
	}

	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()
	return nil
}
