package app

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/controller/v1"
	"github.com/MatheusHenrique129/application-in-go/internal/routes"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/MatheusHenrique129/application-in-go/libraries/logger"
	"github.com/gin-gonic/gin"
)

type App struct {
	appContext       context.Context
	appContextCancel context.CancelFunc
	errorChan        chan error
	LastError        error
	OsSignal         chan os.Signal

	Config *config.Config
	Router *gin.Engine
	DB     *sql.DB
	Logger *util.Logger

	// Controllers
	ProductsController *v1.ProductsController
}

func (a *App) Run() error {
	server := &http.Server{
		Addr:    ":" + a.Config.GetPort(),
		Handler: a.Router,
		// If we don't set these timeouts, we could reach the limit of max open files in the OS. One way to
		// avoid this is to set explicitly the timeouts here. Also, this can fix problems when, for example,
		// a client does not close its connection correctly.
		//
		// See the following issue for more information: https://github.com/golang/go/issues/28272
	}

	a.errorChan = make(chan error)

	go func() {
		a.Logger.InfofWithoutContext("🚀 Starting application on port: %s.", a.Config.GetPort())

		a.errorChan <- server.ListenAndServe()
	}()

	a.waitForShutdown(a.appContext, server)
	a.appContextCancel()

	a.Logger.Infof(a.appContext, "Application stopped.")
	return a.LastError
}

func (a *App) setupDependencies() {
	a.Logger.DebugWithoutContext("Initializing application dependencies...")

	// Controllers
	a.ProductsController = v1.NewProductsController()

	// Services

	// Routes
	r := routes.NewRoutes(a.Config)
	a.Router = r.CreateRouter()

	a.Logger.DebugWithoutContext("Application dependencies initialized.")
}

func (a *App) waitForShutdown(ctx context.Context, srv *http.Server) {
	for {
		select {
		case <-a.OsSignal:
			a.Logger.InfofWithoutContext("Received signal to shutdown the web server. Shutting down in %d milliseconds...", a.Config.GetAppShutdownGraceMilliseconds())
			ctx, cancel := context.WithTimeout(ctx, time.Duration(a.Config.GetAppShutdownGraceMilliseconds())*time.Millisecond)
			defer cancel()

			a.LastError = srv.Shutdown(ctx)
			return
		case err := <-a.errorChan:
			a.Logger.ErrorWithoutContext("Could not start web server", err)

			a.LastError = err
			return
		}
	}
}

func CreateSignalChannel() chan os.Signal {
	exitSignalChannel := make(chan os.Signal, 1)
	signal.Notify(exitSignalChannel, syscall.SIGINT, syscall.SIGTERM)
	return exitSignalChannel
}

func (a *App) Setup() {
	logger.Infof("Setting up configurations and dependencies")
	a.appContext, a.appContextCancel = context.WithCancel(context.Background())

	a.setupDependencies()
	a.Logger.DebugWithoutContext("Application setup finished.")
}

func NewApplication(cfg *config.Config) *App {
	// Previous custom options
	defaultLogger := util.NewLogger("app")

	a := &App{
		Config:   cfg,
		Logger:   defaultLogger,
		OsSignal: CreateSignalChannel(),
	}

	//Load configurations and dependencies
	a.Setup()
	return a
}
