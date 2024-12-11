package controller

import (
	"net/http"
	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"github.com/whatsauth/itmodel"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRestaurant(respw http.ResponseWriter, req *http.Request) {
	var resp itmodel.Response
	resto, err := atdb.GetAllDoc[[]model.Restaurant](config.Mongoconn, "menu_makanan", bson.M{})
	if err != nil {
		resp.Response = err.Error()
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, resto)
}
