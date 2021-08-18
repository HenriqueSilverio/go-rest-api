package health

import "github.com/go-chi/chi"

func Register(router chi.Router) {
	router.Get("/health", index)
}
