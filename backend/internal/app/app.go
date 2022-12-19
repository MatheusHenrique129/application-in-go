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
	"github.com/MatheusHenrique129/application-in-go/internal/controller"
	"github.com/MatheusHenrique129/application-in-go/internal/ginhandler"
	"github.com/MatheusHenrique129/application-in-go/internal/repository"
	"github.com/MatheusHenrique129/application-in-go/internal/routes"
	"github.com/MatheusHenrique129/application-in-go/internal/service"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/MatheusHenrique129/application-in-go/libraries/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type App struct {
	appContext       context.Context
	appContextCancel context.CancelFunc
	osSignal         chan os.Signal
	errorChan        chan error
	LastError        error

	Config     *config.Config
	Router     *gin.Engine
	DB         *sql.DB
	Logger     *util.Logger
	Validator  *validator.Validate
	TimeHelper util.TimeHelper

	// Handlers
	AuthHandler *ginhandler.AuthHandler

	// Repositories
	UserRepository repository.UserRepository

	// Services
	UserService     service.UserService
	ValidateService *service.ValidateService

	// Controllers
	ProductsController *controller.FeedController
	UserController     *controller.UserController
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

	// Validator
	a.Validator = validator.New()
	a.ValidateService = service.NewValidatorService(a.Validator)

	// Repositories
	a.UserRepository = repository.NewUserRepository(a.Config, a.DB)

	// Services
	a.UserService = service.NewUserService(a.Config, a.ValidateService, a.UserRepository)

	// Controllers
	a.ProductsController = controller.NewProductsController()
	a.UserController = controller.NewUserController(a.UserService)

	// Handlers
	a.AuthHandler = ginhandler.NewAuthHandler(a.Config)

	// Routes
	r := routes.NewRoutes(a.Config, a.UserController, a.ProductsController)
	a.Router = r.CreateRouter()

	a.Logger.DebugWithoutContext("Application dependencies initialized.")
}

func (a *App) waitForShutdown(ctx context.Context, srv *http.Server) {
	for {
		select {
		case <-a.osSignal:
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

// CreateSignalChannel Creates a channel that will be notified by signals SIGINT and SIGTERM.
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

func NewApp(db *sql.DB) *App {
	// Previous custom options
	defaultGlobalConfig := config.NewConfig()
	defaultLogger := util.NewLogger("App")

	a := &App{
		Config:     defaultGlobalConfig,
		Logger:     defaultLogger,
		osSignal:   CreateSignalChannel(),
		TimeHelper: util.NewTimeHelper(),
	}

	// TODO rever a necessidade e retirar caso nao precisar futuramente
	if db != nil {
		a.DB = db
	}

	//Load configurations and dependencies
	a.Setup()
	return a
}
