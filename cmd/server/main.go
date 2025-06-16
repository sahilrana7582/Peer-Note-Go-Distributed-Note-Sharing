package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sahil/peernote/internal/config"
	"github.com/sahil/peernote/internal/db"
	"github.com/sahil/peernote/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	db.InitPostgres(cfg.DBURL)
	r := routes.SetupRouter()

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-sigint
		log.Println("🛑 Shutdown signal received. Shutting down server gracefully...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("❌ Server shutdown error: %v", err)
		} else {
			log.Println("✅ Server shutdown completed.")
		}

		close(idleConnsClosed)
	}()

	log.Printf("🚀 Server running on port %s...", cfg.ServerPort)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("❌ Server failed: %v", err)
	}

	<-idleConnsClosed
}
