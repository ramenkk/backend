package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Outlet struct {
    ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Nama          string             `bson:"nama_outlet" json:"nama_outlet"`
    Alamat        string             `bson:"alamat" json:"alamat"`
    KodeOutlet    string             `bson:"kode_outlet" json:"kode_outlet"`
    Barcode       string             `bson:"barcode" json:"barcode"`
    NomorTelepon  string             `bson:"nomor_telepon" json:"nomor_telepon"`
    JamOperasional string            `bson:"jam_operasional" json:"jam_operasional"`
}

type Menu struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    NamaMenu    string             `bson:"nama_menu" json:"nama_menu"`
    Harga       float64            `bson:"harga" json:"harga"`
    Deskripsi   string             `bson:"deskripsi" json:"deskripsi"`
    Gambar      string             `bson:"gambar" json:"gambar"`
    Kategori    string             `bson:"kategori" json:"kategori"`
    OutletID    primitive.ObjectID `bson:"outlet_id,omitempty" json:"outlet_id,omitempty"`
}

type Pesanan struct {
    ID               primitive.ObjectID    `bson:"_id,omitempty" json:"id,omitempty"`
    NamaPelanggan    string                `bson:"nama_pelanggan" json:"nama_pelanggan"`
    OutletID         primitive.ObjectID    `bson:"outlet_id" json:"outlet_id"`
    DaftarMenu       []ItemPesanan         `bson:"daftar_menu" json:"daftar_menu"`
    TotalHarga       float64               `bson:"total_harga" json:"total_harga"`
    StatusPesanan    string                `bson:"status_pesanan" json:"status_pesanan"`
    TanggalPesanan   primitive.DateTime    `bson:"tanggal_pesanan" json:"tanggal_pesanan"`
    Pembayaran       string                `bson:"pembayaran" json:"pembayaran"`
    CatatanPesanan   string                `bson:"catatan_pesanan" json:"catatan_pesanan"`
}

type ItemPesanan struct {
    MenuID       primitive.ObjectID `bson:"menu_id" json:"menu_id"`
    NamaMenu     string             `bson:"nama_menu" json:"nama_menu"`
    Jumlah       int                `bson:"jumlah" json:"jumlah"`
    HargaSatuan  float64            `bson:"harga_satuan" json:"harga_satuan"`
    Subtotal     float64            `bson:"subtotal" json:"subtotal"`
}
