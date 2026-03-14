package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"charm.land/wish/v2"
	cssh "github.com/charmbracelet/ssh"
	"charm.land/wish/v2/activeterm"
	"charm.land/wish/v2/bubbletea"
	"charm.land/wish/v2/logging"

	"github.com/omerdduran/ssh-portfolio/internal/content"
	"github.com/omerdduran/ssh-portfolio/internal/stats"
	"github.com/omerdduran/ssh-portfolio/internal/ui"
)

func statsMiddleware(next cssh.Handler) cssh.Handler {
	return func(s cssh.Session) {
		stats.Global.Connect()
		defer stats.Global.Disconnect()
		next(s)
	}
}

func main() {
	// Warm the content cache before accepting connections
	log.Println("Pre-fetching portfolio content...")
	content.Warm()
	log.Println("Content ready")

	srv, err := wish.NewServer(
		wish.WithAddress("0.0.0.0:23234"),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(ui.TeaHandler),
			activeterm.Middleware(),
			statsMiddleware,
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatalf("Could not create server: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Starting SSH server on 0.0.0.0:23234")
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, cssh.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-done
	log.Println("Shutting down...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
}
