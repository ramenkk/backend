package route

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
	"github.com/gocroot/helper"
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return
	}
	config.SetEnv()

	var method, path string = r.Method, r.URL.Path
	switch {
	case method == "GET" && path == "/data/resto":
		controller.GetRestaurant(w, r)
	case method == "POST" && path == "/data/resto":
		controller.PostRestaurant(w, r) 
	case method == "POST" && helper.URLParam(path, "/webhook/nomor/:nomorwa"):
		controller.PostInboxNomor(w, r)
	// Rute default untuk request yang tidak dikenali.
	default:
		controller.NotFound(w, r) // Mengembalikan response 404 jika rute tidak ditemukan.
	}
}
