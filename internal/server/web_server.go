package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/robovarga/szlh-delegations/internal/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
)

type WebServer struct {
	router *chi.Mux
	games  *repository.GamesRepository
}

func NewWebServer(
	router *chi.Mux,
	healthchecker *HealthCheckHandler,
	gamesRepository *repository.GamesRepository) *WebServer {

	router.Use(middleware.RequestID)

	router.Method(http.MethodGet, "/", healthchecker)
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
	}
}

func (s *WebServer) Serve(ctx context.Context) {
	port := os.Getenv("PORT")
	if port == "" {
		log.Panic(fmt.Errorf("$PORT not set"))
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
			// s.logger.Error(err)
			log.Panic(err)
		}
		close(done)
	}()

	// s.logger.Infof("serving API at port %s", srv.Addr)
	log.Printf("serving API at port %s \n", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// s.logger.Fatal(err)
		log.Panic(err)
	}

	<-done
}
