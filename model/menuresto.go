package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Kategori  string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
	Alamat    string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Telepon   string             `bson:"telepon,omitempty" json:"telepon,omitempty"`
	Status    string             `bson:"status,omitempty" json:"status,omitempty"`
	Waktubuka string             `bson:"waktubuka,omitempty" json:"waktubuka,omitempty"`
	Harga     string             `bson:"harga,omitempty" json:"harga,omitempty"`
	Photo     string             `bson:"photo,omitempty" json:"photo,omitempty"`
}
