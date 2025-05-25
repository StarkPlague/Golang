package main

import "fmt"

type Barang struct {
	Nama_barang string
	Jumlah      int
}

type Gudang struct {
	Daftar_barang []Barang //pakai struct supaya bisa bikin banyak gudang dengan nama yang berbeda
}

func Tambah_barang(gudang *Gudang, barang Barang) {
	gudang.Daftar_barang = append(gudang.Daftar_barang, barang)
}

func Tampilkan_barang(gudang Gudang) {
	for i, v := range gudang.Daftar_barang {
		fmt.Println(i+1, ". ", v.Nama_barang, " | Jumlah: ", v.Jumlah)
	}
}

func Total_barang(gudang Gudang) int {
	total := 0
	for _, v := range gudang.Daftar_barang {
		total += v.Jumlah
	}
	return total
}

//CLI Input gudang
