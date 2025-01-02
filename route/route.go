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
	case method == "GET" && path == "/data/menu_ramen":
		controller.GetMenu_ramen(w, r)
	case method == "POST" && path == "/tambah/menu_ramen":
		controller.Postmenu_ramen(w, r) 
	case method == "POST" && helper.URLParam(path, "/webhook/nomor/:nomorwa"):
		controller.PostInboxNomor(w, r)
	// Rute default untuk request yang tidak dikenali.
	default:
		controller.NotFound(w, r) // Mengembalikan response 404 jika rute tidak ditemukan.
	}
}
