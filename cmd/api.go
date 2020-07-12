package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/kopdar/kopdar-backend/internal/user"
	"github.com/kopdar/kopdar-backend/internal/user/endpoint"
	userhttp "github.com/kopdar/kopdar-backend/internal/user/http"
	"github.com/kopdar/kopdar-backend/internal/user/postgres"
	"github.com/kopdar/kopdar-backend/pkg/pglib"
	_ "github.com/lib/pq"
)

func main() {
	db := sqlx.MustConnect("postgres", "port=5432 user=kevinzola password=ulalala95 dbname=kopdar sslmode=disable")
	err := pglib.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	s := initServer(db)
	s.routes()
	startServer(s)
}

func startServer(s *server) {
	errs := make(chan error, 2)
	go func() {
		port := "9090"
		log.Printf("Listening on :%v ...", port)
		errs <- http.ListenAndServe(":"+port, s)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	log.Println(<-errs)
}

type server struct {
	router     *chi.Mux
	userRouter userhttp.RegisterRoutesFunc
}

func initServer(db *sqlx.DB) *server {
	userRepository := postgres.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userEndpoint := endpoint.New(userService)
	return &server{
		router:     chi.NewRouter(),
		userRouter: userhttp.NewRegisterRoutesFunc(userEndpoint),
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {
	s.router.Route("/v1", func(r chi.Router) {
		s.userRouter(r)
	})
}
