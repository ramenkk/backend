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

		// endpoint outlet
	case method == "GET" && path == "/data/outlet":
		controller.GetOutlet(w, r)
	case method == "GET" && path == "/data/outletbycode":
		controller.GetOutletByCode(w, r)
	case method == "POST" && path == "/tambah/outlet":
		controller.PostOutlet(w, r) 

		// endpoint menu ramen
	case method == "GET" && path == "/data/menu_ramen":
		controller.GetMenu_ramen(w, r)
	case method == "GET" && path == "/data/menu_ramen/byoutletid":
		controller.GetMenuByOutletID(w, r)
	case method == "POST" && path == "/tambah/menu_ramen":
		controller.Postmenu_ramen(w, r) 

		// endpoint pesanan
		controller.GetPesanan(w, r)
	case method == "GET" && path == "/data/pesanan":
		controller.GetPesananByStatus(w, r)
	case method == "GET" && path == "/data/pesananbystatus":

	case method == "POST" && path == "/tambah/pesanan":
		controller.PostPesanan(w, r) 

		 // Endpoint untuk menyelesaikan pesanan
	case method == "PUT" && path == "/complete-order":
		controller.CompleteOrder(w, r)
	case method == "PUT" && path == "/update-order-status":
        controller.UpdateOrderStatus(w, r)
		

		// endpoint item pesanan
		controller.GetItemPesanan(w, r)
	case method == "POST" && path == "/tambah/item_pesanan":
		controller.PostItemPesanan(w, r) 
	case method == "POST" && helper.URLParam(path, "/webhook/nomor/:nomorwa"):
		controller.PostInboxNomor(w, r)
	
	default:
		controller.NotFound(w, r) 
	}
}
