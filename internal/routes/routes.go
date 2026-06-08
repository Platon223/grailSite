package routes

import (
    "net/http"
    "github.com/Platon223/grailSite/internal/domain/index"
)

func Register(mux *http.ServeMux) {
    mux.HandleFunc("GET /", index.Handler)
    // Add more routes here as you grow:
    // mux.HandleFunc("GET /users", users.Handler)
    // mux.HandleFunc("POST /form",  form.Handler)
}
