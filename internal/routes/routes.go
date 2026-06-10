package routes

import (
	"net/http"

	"github.com/Platon223/grailSite/internal/domain/contact"
	"github.com/Platon223/grailSite/internal/domain/docs"
	"github.com/Platon223/grailSite/internal/domain/index"
)

func Register(mux *http.ServeMux) {
    mux.HandleFunc("GET /", index.Handler)
	mux.HandleFunc("GET /docs", docs.Handler)
	mux.HandleFunc("GET /contact", contact.Handler)
}
