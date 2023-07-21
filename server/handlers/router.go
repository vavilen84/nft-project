package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/vavilen84/nft-project/constants"
	"log"
	"net/http"
	"os"
	"time"
)

func MakeHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	v1Router := BuildV1Paths()
	r.Mount("/v1", v1Router)

	return r
}

func BuildV1Paths() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	//r.Route("/posts", func(r chi.Router) {
	//	c := PostsController{}
	//	r.Post("/", c.Create)
	//	r.Put("/{ID}", c.Update)
	//	r.Delete("/{ID}", c.Delete)
	//	r.Get("/{ID}", c.GetOne)
	//	r.Get("/", c.GetAll)
	//})
	r.Route("/security", func(r chi.Router) {
		c := SecurityController{}
		r.Post("/login", c.Login)
		r.Post("/register", c.Register)
	})
	return r
}

func InitHttpServer(handler http.Handler) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)

	return &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: constants.DefaultWriteTimout,
		ReadTimeout:  constants.DefaultReadTimeout,
	}
}
