package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amirabaris/go-auth/internal/config"
	"github.com/amirabaris/go-auth/internal/db"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))

}

func main() {
	// 1. handle config
	cfg := config.LoadConfig()

	pool, err := db.NewPool(context.Background(), cfg)
	if err != nil {
		slog.Error("db", "err", err)
		os.Exit(1)
	}

	defer pool.Close()

	// 2. handle mux
	mux := http.NewServeMux()
	// 3. hande routes
	mux.HandleFunc("GET /hi", handler)

	// 4. config server struct
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	// 5. server.listen in goroutine
	go func() {
		log.Printf("Server listening on %s", cfg.Port)

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	// 6. handle stop
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	// 7. handle context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}

	log.Print("server stopped")
}
