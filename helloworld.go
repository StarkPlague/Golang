package main

import (
	"fmt"
	"strconv"
)

func main() {
	menu := true
	gudang := Gudang{}
	for menu {
		var pilihan int
		fmt.Println("Menu gudang:")
		fmt.Println(`Silahkan pilih menu: 
		1. Tambah Barang
		2. Tampilkan Barang
		3. Totalkan Barang
		4. Keluar`)

		switch pilihan {
		case 1:
			caseOne := true
			var barang string
			var jumlah int
			for caseOne {
				if barang == "n" || strconv.Itoa(jumlah) == "n" {
					caseOne = false
				} else {
					fmt.Println("Masukkan nama barang dan jumlah:")
					fmt.Println("Nama Barang: ")
					fmt.Scan(&barang)
					fmt.Println("Jumlah Barang: ")
					fmt.Scan(&jumlah)
					inputBarang := Barang{barang, jumlah}
					Tambah_barang(&gudang, inputBarang)
					fmt.Println(`ketik 'n' untuk mengakhiri...`)
				}
			}
		case 2:
			Tampilkan_barang(gudang)

		case 3:
			Total_barang(gudang)

		default:
			fmt.Println("app ditutup..")
			menu = false
		}

	}
	RunGudang()
}
