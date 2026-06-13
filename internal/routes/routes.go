package routes

import (
	"net/http"

	"github.com/Platon223/grailSite/internal/domain/contact"
	"github.com/Platon223/grailSite/internal/domain/docs"
	"github.com/Platon223/grailSite/internal/domain/index"
	"github.com/Platon223/grailSite/internal/domain/privacy"
)

func Register(mux *http.ServeMux) {
    mux.HandleFunc("GET /", index.Handler)
	mux.HandleFunc("GET /docs", docs.Handler)
	mux.HandleFunc("GET /contact", contact.Handler)
	mux.HandleFunc("GET /privacy", privacy.Handler)
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}
