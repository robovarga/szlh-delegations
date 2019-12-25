package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/robovarga/szlh-delegations/internal/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
)

type WebServer struct {
	router *chi.Mux
	logger *logrus.Logger
	games  *repository.GamesRepository
}

func NewWebServer(
	router *chi.Mux,
	logger *logrus.Logger,
	healthchecker *HealthCheckHandler,
	listsHandler *ListsHandler,
	gamesHandler *GamesHandler,
	gamesRepository *repository.GamesRepository) *WebServer {

	router.Use(middleware.RequestID, middleware.Heartbeat("/ping"))

	router.Method(http.MethodGet, "/lists", listsHandler)
	router.Method(http.MethodGet, "/games/{id:[0-9]+}", gamesHandler)

	router.Method(http.MethodGet, "/healthz", healthchecker)

	// router.Group(func(r chi.Router) {
	// 	r.Use(chitrc.Middleware(
	// 		chitrc.WithServiceName("notifier.state-updater"),
	// 		chitrc.WithSpanOptions(tracer.Tag("transport", "json")),
	// 	))
	// 	router.Handle("/status/*", updateState)
	// })
	//
	// router.Group(func(r chi.Router) {
	// 	r.Use(chitrc.Middleware(
	// 		chitrc.WithServiceName("notifier.notifications"),
	// 		chitrc.WithSpanOptions(tracer.Tag("transport", "json")),
	// 	))
	// 	r.Use(factory.middlewares()...)
	// 	r.Method(http.MethodPost, "/notification/send", notifier)
	// 	r.Method(http.MethodGet, "/notification/messages", messageStatus)
	// })

	return &WebServer{
		router: router,
		games:  gamesRepository,
		logger: logger,
	}
}

func (s *WebServer) Serve(ctx context.Context) {
	port := os.Getenv("PORT")
	if port == "" {
		s.logger.Fatal("$PORT not set")
	}

	srv := &http.Server{
		Addr:        ":" + port,
		Handler:     s.router,
		ReadTimeout: 30 * time.Second,
	}

	done := make(chan struct{})

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			s.logger.Error(err)
		}
		close(done)
	}()

	s.logger.Infof("serving API at port %s", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		s.logger.Fatal(err)
	}

	<-done
}
