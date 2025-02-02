package route

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
	"github.com/gocroot/handler"
	"github.com/gocroot/helper"
	"github.com/gocroot/middleware"
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return
	}
	config.SetEnv()

	var method, path string = r.Method, r.URL.Path
	switch {


		// endpoint menu ramen
	case method == "GET" && path == "/data/menu_ramen":
		controller.GetMenu_ramen(w, r)
	case method == "POST" && path == "/tambah/menu_ramen":
		middleware.CSRFMiddleware(http.HandlerFunc(controller.Postmenu_ramen)).ServeHTTP(w, r)
	case method == "PUT" && path == "/ubah/menu_ramen":
		controller.PutMenu(w, r)
	case method == "DELETE" && path == "/hapus/menu_ramen":
		controller.DeleteMenu(w, r)

	case method == "GET" && path == "/csrf-token":
		handler.CSRFToken(w, r) // Endpoint untuk menghasilkan token CSRF

		// endpoint pesanan
	case method == "GET" && path == "/data/pesanan":
		controller.GetPesanan(w, r)
	case method == "GET" && path == "/data/byid":
		controller.GetPesananByID(w, r)
	
	
	
	case method == "GET" && path == "/data/bystatus":
		controller.GetPesananByStatus(w, r)
	case method == "POST" && path == "/tambah/pesanan":
		controller.PostPesanan(w, r)
	case method == "PATCH" && path == "/update/status":
		controller.UpdatePesananStatus(w, r)

		// endpoint item pesanan
		controller.GetItemPesanan(w, r)
	case method == "POST" && path == "/tambah/item_pesanan":
		controller.PostItemPesanan(w, r)
	case method == "POST" && helper.URLParam(path, "/webhook/nomor/:nomorwa"):
		controller.PostInboxNomor(w, r)

		// Rute untuk admin (login, logout, register, dashboard, aktivitas).
	case method == "POST" && path == "/admin/login":
		handler.Login(w, r) // Login admin.
	case method == "GET" && path == "/data/activity":
		controller.GetActivity(w, r)
	case method == "POST" && path == "/admin/logout":
		handler.Logout(w, r) // Logout admin.
	case method == "POST" && path == "/admin/register":
		handler.RegisterAdmin(w, r) // Registrasi admin baru.
	case method == "GET" && path == "/admin/dashboard":
		// Middleware autentikasi untuk dashboard admin.
		middleware.AuthMiddleware(http.HandlerFunc(handler.DashboardAdmin)).ServeHTTP(w, r)


	default:
		controller.NotFound(w, r)
	}
}



