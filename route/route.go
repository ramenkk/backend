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

	// 🔹 Endpoint Menu Ramen
	case method == "GET" && path == "/data/menu_ramen":
		controller.GetMenu_ramen(w, r)
	case method == "POST" && path == "/tambah/menu_ramen":
		middleware.CSRFMiddleware(http.HandlerFunc(controller.Postmenu_ramen)).ServeHTTP(w, r)
	case method == "PUT" && path == "/ubah/menu_ramen":
		middleware.CSRFMiddleware(http.HandlerFunc(controller.PutMenu)).ServeHTTP(w, r)
	case method == "DELETE" && path == "/hapus/menu_ramen":
		middleware.CSRFMiddleware(http.HandlerFunc(controller.DeleteMenu)).ServeHTTP(w, r)

	// 🔹 Endpoint Pesanan
	case method == "GET" && path == "/data/pesanan":
		controller.GetPesanan(w, r)
	case method == "GET" && path == "/data/byid":
		controller.GetPesananByID(w, r)
	case method == "GET" && path == "/data/bystatus":
		controller.GetPesananByStatus(w, r)
	case method == "POST" && path == "/tambah/pesanan":
		middleware.CSRFMiddleware(http.HandlerFunc(controller.PostPesanan)).ServeHTTP(w, r)
	case method == "PATCH" && path == "/update/status":
		middleware.CSRFMiddleware(http.HandlerFunc(controller.UpdatePesananStatus)).ServeHTTP(w, r)

	// 🔹 Endpoint Item Pesanan (Fix Rute yang Hilang)
	case method == "GET" && path == "/data/item_pesanan":
		controller.GetItemPesanan(w, r)
	case method == "POST" && path == "/tambah/item_pesanan":
		middleware.CSRFMiddleware(http.HandlerFunc(controller.PostItemPesanan)).ServeHTTP(w, r)

	// 🔹 Webhook (Nomor WhatsApp)
	case method == "POST" && helper.URLParam(path, "/webhook/nomor/:nomorwa"):
		controller.PostInboxNomor(w, r)

	// 🔹 Rute untuk Admin (Login, Logout, Register, Dashboard, Aktivitas)
	case method == "POST" && path == "/admin/login":
		handler.Login(w, r)
	case method == "GET" && path == "/data/activity":
		middleware.AuthMiddleware(http.HandlerFunc(controller.GetActivity)).ServeHTTP(w, r)
	case method == "POST" && path == "/admin/logout":
		middleware.AuthMiddleware(http.HandlerFunc(handler.Logout)).ServeHTTP(w, r)
	case method == "POST" && path == "/admin/register":
		middleware.AuthMiddleware(http.HandlerFunc(handler.RegisterAdmin)).ServeHTTP(w, r)
	case method == "GET" && path == "/admin/dashboard":
		middleware.AuthMiddleware(http.HandlerFunc(handler.DashboardAdmin)).ServeHTTP(w, r)

	// 🔹 Not Found (Rute tidak dikenal)
	default:
		controller.NotFound(w, r)
	}
}
