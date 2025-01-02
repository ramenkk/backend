package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"github.com/whatsauth/itmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
