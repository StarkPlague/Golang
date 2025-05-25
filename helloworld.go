package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	menu := true
	gudang := Gudang{}
	for menu {
		var pilihan int
		fmt.Println("\nMenu gudang:")
		fmt.Println(`Silahkan pilih menu: 
		1. Tambah Barang
		2. Tampilkan Barang
		3. Totalkan Barang
		4. Keluar`)

		fmt.Scan(&pilihan)
		bufio.NewReader(os.Stdin).ReadString('\n') // <-- buang newline sisa

		switch pilihan {
		case 1:
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Println("Masukkan data barang (ketik 'n' untuk mengakhiri)...")

				fmt.Print("Nama Barang: ")
				nama, _ := reader.ReadString('\n')
				nama = strings.TrimSpace(nama)

				if nama == "n" {
					break
				}

				fmt.Print("Jumlah Barang: ")
				jumlahStr, _ := reader.ReadString('\n')
				jumlahStr = strings.TrimSpace(jumlahStr)

				jumlah, err := strconv.Atoi(jumlahStr)
				if err != nil {
					fmt.Println("Jumlah harus berupa angka. Coba lagi.")
					continue
				}

				barangBaru := Barang{Nama_barang: nama, Jumlah: jumlah}
				Tambah_barang(&gudang, barangBaru)

				fmt.Print("Barang berhasil ditambahkan!\n")
			}

		case 2:
			Tampilkan_barang(gudang)

		case 3:
			total := Total_barang(gudang)
			fmt.Println("Total barang: ", total)

		default:
			fmt.Println("app ditutup..")
			menu = false
		}

	}
}
