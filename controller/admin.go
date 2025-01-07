package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/module"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if config.Mongoconn == nil {
		http.Error(w, "Database connection not established", http.StatusInternalServerError)
		return
	}

	// Access the admins collection
	collection := config.Mongoconn.Collection("admins")

	// Find admin by username
	admin, err := module.FindAdminByUsername(collection, req.Username)
	if err != nil || !helper.CheckPasswordHash(req.Password, admin.Password) || admin.Role != "admin" {
		http.Error(w, "Invalid credentials or unauthorized", http.StatusUnauthorized)
		return
	}

	// Generate token for the admin
	token, err := helper.GenerateToken(admin.Username, admin.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond with the token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
