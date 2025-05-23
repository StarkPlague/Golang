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
		fmt.Println("No. ", i+1, v.Nama_barang, " | Jumlah: ", v.Jumlah)
	}
}

func Total_barang(gudang Gudang) int {
	total := 0
	for _, v := range gudang.Daftar_barang {
		total += v.Jumlah
	}
	return total
}

func RunGudang() {
	gudang := Gudang{}
	barang1 := Barang{"Buku", 10}
	barang2 := Barang{"Pensil", 5}
	Tambah_barang(&gudang, barang1)
	Tambah_barang(&gudang, barang2)

	fmt.Println("Daftar Barang di Gudang:")
	Tampilkan_barang(gudang)
}
