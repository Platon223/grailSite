package index


import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "internal/domain/index/templates/index.html")
}
