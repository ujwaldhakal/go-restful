package main

import (
	"flag"
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	_ "github.com/lib/pq"
	"github.com/user/sites/app/internal/forecast"
	"github.com/user/sites/app/internal/webhook"
	"github.com/user/sites/app/internal/city"
	"github.com/user/sites/app/internal/temperature"
	"github.com/user/sites/app/internal/config"
	"github.com/user/sites/app/internal/errors"
	"github.com/user/sites/app/pkg/accesslog"
	"github.com/user/sites/app/pkg/dbcontext"
	"github.com/user/sites/app/pkg/graceful"
	"github.com/user/sites/app/pkg/log"
	"net/http"
	"os"
	"time"

)

// Version indicates the current version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

func main() {
	flag.Parse()
	// create root logger tagged with server version
	logger := log.New().With(nil, "version", Version)

	// load application configurations
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	// connect to the database
	db, err := dbx.MustOpen("postgres", cfg.DSN)
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}

	db.QueryLogFunc = log.DBQuery(logger)
	db.ExecLogFunc = log.DBExec(logger)
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error(err)
		}
	}()

	// build HTTP server
	address := fmt.Sprintf(":%v", cfg.ServerPort)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(logger, dbcontext.New(db), cfg),
	}

	// start the HTTP server with graceful shutdown
	go graceful.Shutdown(hs, 10*time.Second, logger)
	logger.Infof("server %v is running at %v", Version, address)
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err)
		os.Exit(-1)
	}
}

func buildHandler(logger log.Logger, db *dbcontext.DB, cfg *config.Config) http.Handler {
	router := routing.New()

	router.Use(
		accesslog.Handler(logger),
		errors.Handler(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
	)



	rg := router.Group("")


	city.RegisterHandlers(rg.Group(""),
		city.NewService(city.NewRepository(db, logger), logger),
		logger,
	)

	temperature.RegisterHandlers(rg.Group(""),
		temperature.NewService(temperature.NewRepository(db, logger), logger),
		logger,
	)

	forecast.RegisterHandlers(rg.Group(""),
		forecast.NewService(forecast.NewRepository(db, logger), logger),
		logger,
	)

	webhook.RegisterHandlers(rg.Group(""),
    webhook.NewService(webhook.NewRepository(db, logger), logger),
		logger,
	)


	return router
}
