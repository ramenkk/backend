package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"github.com/whatsauth/itmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetOutlet(respw http.ResponseWriter, req *http.Request) {
	var resp itmodel.Response
	outlets, err := atdb.GetAllDoc[[]model.Outlet](config.Mongoconn, "outlet", bson.M{})
	if err != nil {
		resp.Response = err.Error()
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, outlets)
}

func ValidateKodeOutlet(respw http.ResponseWriter, req *http.Request) {
    kodeOutlet := req.URL.Query().Get("kode_outlet")
    if kodeOutlet == "" {
        respondWithError(respw, http.StatusBadRequest, "Kode outlet harus disertakan")
        return
    }

    filter := bson.M{"kode_outlet": kodeOutlet}

    var outlet model.Outlet
    err := config.Mongoconn.Collection("outlet").FindOne(context.Background(), filter).Decode(&outlet)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            respondWithError(respw, http.StatusNotFound, "Kode outlet tidak ditemukan")
        } else {
            respondWithError(respw, http.StatusInternalServerError, fmt.Sprintf("Terjadi kesalahan: %v", err))
        }
        return
    }

    respw.Header().Set("Content-Type", "application/json")
    respw.WriteHeader(http.StatusOK)
    json.NewEncoder(respw).Encode(map[string]interface{}{
        "status":     "success",
        "outlet_id":  outlet.ID.Hex(),
		"outlet_nama": outlet.Nama, 
    })
}

func GetOutletByCode(respw http.ResponseWriter, req *http.Request) {

	kodeOutlet := req.URL.Query().Get("kode_outlet")
	if kodeOutlet == "" {

		http.Error(respw, "Kode outlet harus disertakan", http.StatusBadRequest)
		return
	}

	// Membuat filter untuk mencari outlet berdasarkan kode_outlet
	filter := bson.M{"kode_outlet": kodeOutlet}

	var outlet model.Outlet
	outlet, err := atdb.GetOneDoc[model.Outlet](config.Mongoconn, "outlet", filter)
	if err != nil {

		if err == mongo.ErrNoDocuments {
			http.Error(respw, "Outlet tidak ditemukan", http.StatusNotFound)
		} else {

			http.Error(respw, fmt.Sprintf("Terjadi kesalahan: %v", err), http.StatusInternalServerError)
		}
		return
	}

	// Mengatur header respons untuk mengirimkan data dalam format JSON
	respw.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(respw).Encode(outlet); err != nil {

		http.Error(respw, fmt.Sprintf("Terjadi kesalahan saat mengirimkan data: %v", err), http.StatusInternalServerError)
	}
}

func PostOutlet(respw http.ResponseWriter, req *http.Request) {
	var outlet model.Outlet
	if err := json.NewDecoder(req.Body).Decode(&outlet); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, itmodel.Response{Response: err.Error()})
		return
	}

	result, err := config.Mongoconn.Collection("outlet").InsertOne(context.Background(), outlet)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, itmodel.Response{Response: err.Error()})
		return
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	helper.WriteJSON(respw, http.StatusOK, itmodel.Response{Response: fmt.Sprintf("Outlet berhasil disimpan dengan ID: %s", insertedID.Hex())})
}

func GetMenu_ramen(respw http.ResponseWriter, req *http.Request) {
	var resp itmodel.Response
	resto, err := atdb.GetAllDoc[[]model.Menu](config.Mongoconn, "menu_ramen", bson.M{})
	if err != nil {
		resp.Response = err.Error()
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, resto)
}
func GetMenuByOutletID(respw http.ResponseWriter, req *http.Request) {
    outletID := req.URL.Query().Get("outlet_id")
    if outletID == "" {
        respondWithError(respw, http.StatusBadRequest, "Outlet ID harus disertakan")
        return
    }

    objID, err := primitive.ObjectIDFromHex(outletID)
    if err != nil {
        respondWithError(respw, http.StatusBadRequest, "Outlet ID tidak valid")
        return
    }

    filter := bson.M{"outlet_id": objID}

    var menu []model.Menu
    menu, err = atdb.GetFilteredDocs[[]model.Menu](config.Mongoconn, "menu_ramen", filter, nil)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            respondWithError(respw, http.StatusNotFound, "Menu tidak ditemukan untuk outlet ini")
        } else {
            respondWithError(respw, http.StatusInternalServerError, fmt.Sprintf("Terjadi kesalahan: %v", err))
        }
        return
    }

    // Return response JSON
    respw.Header().Set("Content-Type", "application/json")
    respw.WriteHeader(http.StatusOK)
    json.NewEncoder(respw).Encode(map[string]interface{}{
        "status": "success",
        "data":   menu,
    })
}

// Helper function to respond with error
func respondWithError(respw http.ResponseWriter, code int, message string) {
    respw.Header().Set("Content-Type", "application/json")
    respw.WriteHeader(code)
    json.NewEncoder(respw).Encode(map[string]string{"error": message})
}


func Postmenu_ramen(respw http.ResponseWriter, req *http.Request) {

	var restoran model.Menu
	if err := json.NewDecoder(req.Body).Decode(&restoran); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, itmodel.Response{Response: err.Error()})
		return
	}

	result, err := config.Mongoconn.Collection("menu_ramen").InsertOne(context.Background(), restoran)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, itmodel.Response{Response: err.Error()})
		return
	}

	insertedID := result.InsertedID.(primitive.ObjectID)

	helper.WriteJSON(respw, http.StatusOK, itmodel.Response{Response: fmt.Sprintf("Menu ramen berhasil disimpan dengan ID: %s", insertedID.Hex())})
}

func GetPesanan(respw http.ResponseWriter, req *http.Request) {
	var resp itmodel.Response
	orders, err := atdb.GetAllDoc[[]model.Pesanan](config.Mongoconn, "pesanan", bson.M{})
	if err != nil {
		resp.Response = err.Error()
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, orders)
}

func GetPesananByOutletID(respw http.ResponseWriter, req *http.Request) {
    outletID := req.URL.Query().Get("outlet_id")
    if outletID == "" {
        respondWithError(respw, http.StatusBadRequest, "Outlet ID harus disertakan")
        return
    }

    objID, err := primitive.ObjectIDFromHex(outletID)
    if err != nil {
        respondWithError(respw, http.StatusBadRequest, "Outlet ID tidak valid")
        return
    }

    filter := bson.M{"outlet_id": objID}

    var pesanan []model.Pesanan
    pesanan, err = atdb.GetFilteredDocs[[]model.Pesanan](config.Mongoconn, "pesanan", filter, nil)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            respondWithError(respw, http.StatusNotFound, "Pesanan tidak ditemukan untuk outlet ini")
        } else {
            respondWithError(respw, http.StatusInternalServerError, fmt.Sprintf("Terjadi kesalahan: %v", err))
        }
        return
    }

    // Return response JSON
    respw.Header().Set("Content-Type", "application/json")
    respw.WriteHeader(http.StatusOK)
    json.NewEncoder(respw).Encode(map[string]interface{}{
        "status": "success",
        "data":   pesanan,
    })
}

func GetPesananByStatus(respw http.ResponseWriter, req *http.Request) {

	status := req.URL.Query().Get("status")
	if status == "" {
		http.Error(respw, "Status pesanan harus disertakan", http.StatusBadRequest)
		return
	}

	filter := bson.M{"status": status}

	var pesanan []model.Pesanan
	pesanan, err := atdb.GetFilteredDocs[[]model.Pesanan](config.Mongoconn, "pesanan", filter, nil)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(respw, "Pesanan tidak ditemukan dengan status ini", http.StatusNotFound)
		} else {
			http.Error(respw, fmt.Sprintf("Terjadi kesalahan: %v", err), http.StatusInternalServerError)
		}
		return
	}

	helper.WriteJSON(respw, http.StatusOK, pesanan)
}

func PostPesanan(respw http.ResponseWriter, req *http.Request) {
    var pesanan model.Pesanan

    // Decode request body into pesanan
    if err := json.NewDecoder(req.Body).Decode(&pesanan); err != nil {
        log.Println("Error decoding request body:", err)
        helper.WriteJSON(respw, http.StatusBadRequest, itmodel.Response{Response: "Invalid request body"})
        return
    }

    if len(pesanan.DaftarMenu) == 0 {
        helper.WriteJSON(respw, http.StatusBadRequest, itmodel.Response{Response: "Daftar menu cannot be empty"})
        return
    }

 
    pesanan.StatusPesanan = "Baru"
    pesanan.Pembayaran = "Cash"
    pesanan.TanggalPesanan = primitive.NewDateTimeFromTime(time.Now())


    log.Println("Pesanan data received:", pesanan)

 
    result, err := config.Mongoconn.Collection("pesanan").InsertOne(context.Background(), pesanan)
    if err != nil {
        log.Println("Error inserting pesanan:", err)
        helper.WriteJSON(respw, http.StatusInternalServerError, itmodel.Response{Response: "Failed to insert pesanan"})
        return
    }

    insertedID, ok := result.InsertedID.(primitive.ObjectID)
    if !ok {
        log.Println("Error asserting inserted ID to ObjectID")
        helper.WriteJSON(respw, http.StatusInternalServerError, itmodel.Response{Response: "Failed to process inserted ID"})
        return
    }

    log.Println("Pesanan berhasil disimpan dengan ID:", insertedID.Hex())

    helper.WriteJSON(respw, http.StatusOK, itmodel.Response{Response: fmt.Sprintf("Pesanan berhasil disimpan dengan ID: %s", insertedID.Hex())})
}



func GetItemPesanan(respw http.ResponseWriter, req *http.Request) {
	var resp itmodel.Response
	items, err := atdb.GetAllDoc[[]model.ItemPesanan](config.Mongoconn, "item_pesanan", bson.M{})
	if err != nil {
		resp.Response = err.Error()
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, items)
}

func PostItemPesanan(respw http.ResponseWriter, req *http.Request) {
	var item model.ItemPesanan
	if err := json.NewDecoder(req.Body).Decode(&item); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, itmodel.Response{Response: err.Error()})
		return
	}

	result, err := config.Mongoconn.Collection("item_pesanan").InsertOne(context.Background(), item)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, itmodel.Response{Response: err.Error()})
		return
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	helper.WriteJSON(respw, http.StatusOK, itmodel.Response{Response: fmt.Sprintf("Item pesanan berhasil disimpan dengan ID: %s", insertedID.Hex())})
}

func CompleteOrder(respw http.ResponseWriter, req *http.Request) {

	orderID := req.URL.Query().Get("order_id")
	if orderID == "" {
		http.Error(respw, "Order ID harus disertakan", http.StatusBadRequest)
		return
	}

	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		http.Error(respw, "Order ID tidak valid", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": objID}

	update := bson.M{
		"$set": bson.M{
			"status_pesanan": "Selesai",
			"waktu_terima":   primitive.NewDateTimeFromTime(time.Now()),
		},
	}

	result, err := config.Mongoconn.Collection("pesanan").UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(respw, fmt.Sprintf("Terjadi kesalahan: %v", err), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(respw, "Pesanan tidak ditemukan", http.StatusNotFound)
		return
	}

	helper.WriteJSON(respw, http.StatusOK, itmodel.Response{Response: "Pesanan berhasil diselesaikan"})
}

func UpdateOrderStatus(respw http.ResponseWriter, req *http.Request) {
	orderID := req.URL.Query().Get("order_id")
	status := req.URL.Query().Get("status")

	if orderID == "" || status == "" {
		http.Error(respw, "Order ID dan status harus disertakan", http.StatusBadRequest)
		return
	}

	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		http.Error(respw, "Order ID tidak valid", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"status_pesanan": status}}

	result, err := config.Mongoconn.Collection("pesanan").UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(respw, fmt.Sprintf("Terjadi kesalahan: %v", err), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(respw, "Pesanan tidak ditemukan", http.StatusNotFound)
		return
	}

	helper.WriteJSON(respw, http.StatusOK, itmodel.Response{Response: "Status pesanan berhasil diperbarui"})
}
