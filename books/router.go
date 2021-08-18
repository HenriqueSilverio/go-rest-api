package books

import (
	"github.com/go-chi/chi"
)

func Register(router chi.Router) {
	router.Route("/books", func(r chi.Router) {
		r.Get("/", index)
		r.Get("/{bookID}", show)
	})
}
