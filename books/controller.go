package books

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server hits: /books")
}

func show(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "bookID")
	fmt.Fprintf(w, "Server hits: /books/%s", bookID)
}
