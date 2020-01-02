package server

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/robovarga/szlh-delegations/internal/handler"
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
	healthchecker *handler.HealthCheckHandler,
	listsHandler *handler.ListsHandler,
	gamesHandler *handler.GamesHandler,
	refsHandler *handler.RefereesHandler,
	gamesRepository *repository.GamesRepository) *WebServer {

	router.Use(middleware.RequestID)

	router.Method(http.MethodGet, "/healthz", healthchecker)

	router.Method(http.MethodGet, "/lists", listsHandler)

	router.Route("/games", func(r chi.Router) {
		r.Get("/{listId:[0-9]+}", gamesHandler.GetByListID)

		// r.Get("/{refId:[0-9]+}", refsHandler.GetReferee)
	})

	router.Route("/referees", func(r chi.Router) {
		r.Get("/", refsHandler.GetAll)

		r.Get("/{refId:[0-9]+}", refsHandler.GetReferee)
	})

	return &WebServer{
		router: router,
		games:  gamesRepository,
		logger: logger,
	}
}

func (s *WebServer) Serve(ctx context.Context) {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	s.logger.Fatal("$PORT not set")
	// }

	port := "8080"

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
