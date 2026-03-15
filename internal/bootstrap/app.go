package bootstrap

import (
	"HGoComicMosaic/internal/config"
	"HGoComicMosaic/internal/platform/database"
	httptransport "HGoComicMosaic/internal/transport/http"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Config     *config.Config
	DB         *database.Client
	httpServer *httptransport.Server
}

func NewApp() (*App, error) {
	cfg, err := config.Load()

	if err != nil {
		return nil, err
	}

	db, err := database.New(context.Background(), cfg.Database.DSN)
	if err != nil {
		return nil, err
	}

	r := httptransport.NewRouter()
	s := httptransport.NewServer(cfg.HTTP.Port, r)
	fmt.Printf("当前启动的端口 %d", cfg.HTTP.Port)

	return &App{
		Config:     cfg,
		DB:         db,
		httpServer: s,
	}, nil
}

func (a *App) Addr() string {
	return fmt.Sprintf("0.0.0.0:%d", a.Config.HTTP.Port)
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		if err := a.httpServer.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		return a.shutdown()
	case <-sigCh:
		return a.shutdown()
	case err := <-errCh:
		return err
	}

}

func (a *App) shutdown() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	if err := a.DB.Close(); err != nil {
		return err
	}

	return nil
}
