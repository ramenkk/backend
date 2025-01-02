package controller

import (
	"context"
	"encoding/json"
	"fmt"
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

func PostPesanan(respw http.ResponseWriter, req *http.Request) {
	var pesanan model.Pesanan
	if err := json.NewDecoder(req.Body).Decode(&pesanan); err != nil {
		helper.WriteJSON(respw, http.StatusBadRequest, itmodel.Response{Response: err.Error()})
		return
	}
	pesanan.StatusPesanan = "Baru"
	pesanan.TanggalPesanan = primitive.NewDateTimeFromTime(time.Now())

	result, err := config.Mongoconn.Collection("pesanan").InsertOne(context.Background(), pesanan)
	if err != nil {
		helper.WriteJSON(respw, http.StatusInternalServerError, itmodel.Response{Response: err.Error()})
		return
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
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
