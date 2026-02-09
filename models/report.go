package models

type ProdukTerjual struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

type Report struct {
	TotalRevenue   int           `json:"total_revenue"`
	TotalTransaksi int           `json:"total_transaksi"`
	ProdukTerlaris ProdukTerjual `json:"produk_terlaris"`
}
