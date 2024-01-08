package api

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"applicationDesignTest/api/middleware"
	"applicationDesignTest/api/v1/orders/handlers"
	"applicationDesignTest/internals/domain/usecases"
	"github.com/gorilla/mux"
	"golang.org/x/sync/semaphore"
)

const (
	OrdersPrefix = "/orders"
	MaxRateLimit = 100
)

type WebApp struct {
	ordersHandlers *handlers.OrderHandlers
	router         *mux.Router
	logger         *log.Logger
}

func NewWebApp(
	ordersUseCases usecases.OrdersUseCases,
	logger *log.Logger,
) *WebApp {

	app := WebApp{
		ordersHandlers: handlers.NewOrderHandlers(ordersUseCases, logger),
		router:         mux.NewRouter(),
		logger:         logger,
	}

	connectionsSem := semaphore.NewWeighted(MaxRateLimit)
	limiterMW := middleware.RateLimit{ConnectionsSem: connectionsSem}.Next

	app.router.Use(limiterMW)
	app.router.Use(middleware.WriteMetrics)

	ordersPath := OrdersPrefix
	userOrdersPath := ordersPath + "/{id}/"

	app.router.
		HandleFunc(ordersPath, app.ordersHandlers.Post).
		Methods(http.MethodPost).
		Name("order")

	app.router.
		HandleFunc(userOrdersPath, app.ordersHandlers.GetByUserId).
		Methods(http.MethodGet).
		Name("user-orders")

	return &app
}

func (a *WebApp) Run() {
	server := &http.Server{
		Handler: a.router,
		Addr:    ":8080",
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			a.logger.Println(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	<-ctx.Done()

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxShutdown); err != nil {
		a.logger.Printf("shutdown error: %s", err.Error())
	}
}
